function Install-AppStreamFile {
  param(
    [Parameter(Mandatory=$true, Position=0)]
    [string]$Version
  )

  $repo    = "aslamcodes/appstreamfile"
  $baseUrl = "https://github.com/$repo/releases/download/$Version"
  $dir     = "$env:TEMP\appstreamfile"

  if (!(Test-Path $dir)) {
    New-Item -ItemType Directory -Path $dir | Out-Null
  }

  $arch = switch ($env:PROCESSOR_ARCHITECTURE) {
    "AMD64" { "x86_64" }
    "x86"   { "i386" }
    "ARM64" { "arm64" }
    default { throw "Unsupported architecture" }
  }

  $file = "appstreamfile_Windows_$arch.zip"
  $url  = "$baseUrl/$file"
  $zip  = Join-Path $dir $file

  Write-Host "Downloading $file..."
  Invoke-WebRequest $url -OutFile $zip
  Expand-Archive $zip $dir -Force
  Remove-Item $zip

  $exe = Join-Path $dir "appstreamfile.exe"
  if (!(Test-Path $exe)) { throw "Binary not found" }

  Write-Host "Installation complete!"
  & $exe --help
}
