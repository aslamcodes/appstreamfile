# --- AppStreamFile runtime downloader for Windows ---

$repo    = "aslamcodes/appstreamfile"
$version = "v0.1.0"
$baseUrl = "https://github.com/$repo/releases/download/$version"
$dir     = "$env:TEMP\appstreamfile"

if (!(Test-Path $dir)) { New-Item -ItemType Directory -Path $dir | Out-Null }

# Detect architecture from Windows environment
$pa   = $env:PROCESSOR_ARCHITECTURE
$arch = switch ($pa) {
  "AMD64"      { "x86_64" }
  "x86"        { "i386" }
  "ARM64"      { "arm64" }
  default      { throw "Unsupported architecture: $pa" }
}

$file = "appstreamfile_Windows_$arch.zip"
$url  = "$baseUrl/$file"

Write-Host "Downloading $file..."
$zipPath = Join-Path $dir $file
Invoke-WebRequest -Uri $url -OutFile $zipPath

Write-Host "Extracting..."
Expand-Archive -Path $zipPath -DestinationPath $dir -Force
Remove-Item $zipPath

$exe = Join-Path $dir "appstreamfile.exe"
if (!(Test-Path $exe)) { throw "Binary not found after extraction" }

Write-Host "Binary ready at: $exe"
Write-Host "Running --help"
& $exe --help
