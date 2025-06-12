# Configuration
$repoOwner = "kazxieo"
$repoName = "heartbeats"
$binaryName = "heartbeats.exe" # Replace with the actual binary name
$assetExtension = ".exe" # Use ".exe" if it's a direct binary

# GitHub API URL for latest release
$releaseUrl = "https://api.github.com/repos/$repoOwner/$repoName/releases/latest"

# Temporary working directory
$tempDir = "$env:TEMP\$repoName"
New-Item -ItemType Directory -Force -Path $tempDir | Out-Null
Set-Location $tempDir

# Get latest release data
Write-Host "Fetching latest release info..."
$releaseData = Invoke-RestMethod -Uri $releaseUrl -Headers @{ "User-Agent" = "Mozilla/5.0" }

# Find the asset download URL
$asset = $releaseData.assets | Where-Object { $_.name -like "*$assetExtension" } | Select-Object -First 1
if (-not $asset) {
    Write-Error "No asset with extension '$assetExtension' found in the latest release."
    exit 1
}

$downloadUrl = $asset.browser_download_url
$assetPath = "$tempDir\$($asset.name)"

# Download the asset
Write-Host "Downloading asset: $($asset.name)..."
Invoke-WebRequest -Uri $downloadUrl -OutFile $assetPath

# Extract if it's a ZIP file
if ($assetExtension -eq ".zip") {
    Write-Host "Extracting zip archive..."
    Expand-Archive -Path $assetPath -DestinationPath $tempDir -Force
}

# Find the binary
$binaryPath = Get-ChildItem -Path $tempDir -Recurse -Filter $binaryName | Select-Object -First 1
if (-not $binaryPath) {
    Write-Error "Binary '$binaryName' not found after extraction."
    exit 1
}

# Get Windows Startup folder
$startupFolder = [Environment]::GetFolderPath("Startup")
$destPath = Join-Path -Path $startupFolder -ChildPath $binaryName

# Copy binary to Startup
Write-Host "Copying binary to Startup folder..."
Copy-Item -Path $binaryPath.FullName -Destination $destPath -Force

Write-Host "✅ Installed '$binaryName' to Startup successfully."
