# Appstreamfile
A Dockerfile for AWS Image Builders Built in Go. A declarative spec for AWS AppStream images. Appstreamfile defines everything that goes into an image, while your automation system handles execution. 

# Why
AWS AppStream provides a CLI, but it only handles high-level image operations like updating stack applications or building images. The real pain is preparing the image itself. Its the installers, custom scripts, environment tweaks, catalog entries, and session scripts. Doing this by hand or via plain scripting is messy and impossible to audit cleanly.

**Appstreamfile** is the missing piece. It acts like a Dockerfile for AppStream image builders. Install it inside the image builder instance with your favorite automation of choice, point it to a config file, and execute. Your image becomes reproducible, reviewable, and predictable.

# Features
1. Single config file describing the image
2. Scripted installers
3. App catalog configuration
4. System/user session scripts
5. File provisioning
6. Declarative image metadata

# Planned Features
1. Registry Configs
2. Scaffolded startup Apps - Customise your app startup
3. S3 and Git backend support
4. State persistence within image builders 

# Installation
## Windows Installation
You can install and run **AppStreamFile** without cloning the repository.  
The script detects your system architecture and downloads the correct binary automatically.

```powershell
iwr https://raw.githubusercontent.com/aslamcodes/appstreamfile/main/scripts/install.ps1 | iex
```

## Linux
***Supported in future releases***

# Usage
Use **local** backend source when the config file is within image builders. S3, Git Support is WIP
```sh
    appstreamfile -source local -location config.yaml
```

# Sample config
```yaml
platform: "windows"
installers:
  - executable: "powershell"
    installScript: |
      Write-Host "Hello World"

  - executable: "powershell"
    installScript: |
      echo "Setting up environment"

catalog:
  - name: "Notepad"
    path: "C:\\Windows\\System32\\notepad.exe"
    display_name: "Notepad"
    parameters: ""
    icon_path: "C:\\Windows\\System32\\notepad.exe"
    working_dir: "C:\\Windows\\System32"

files:
  - path: "C:\\AppStream\\Scripts\\Start-System.ps1"
    content: |
      Write-EventLog -LogName Application -Source AppStream -EventID 100 -Message "System session start"

  - path: "C:\\AppStream\\Scripts\\Start-User.ps1"
    content: |
      Write-Host "User profile initialization"

  - path: "C:\\AppStream\\Scripts\\End-System.ps1"
    content: |
      Write-EventLog -LogName Application -Source AppStream -EventID 200 -Message "System cleanup"

  - path: "C:\\AppStream\\Scripts\\End-User.ps1"
    content: |
      Write-Host "User session cleanup"

session_scripts:
  session_start:
    executables:
      - context: "system"
        filename: "C:\\AppStream\\Scripts\\Start-System.ps1"
        arguments: "-Init"
        s3LogEnabled: true
      - context: "user"
        filename: "C:\\AppStream\\Scripts\\Start-User.ps1"
        arguments: "-UserSetup"
        s3LogEnabled: true
    waitingTime: 60

  session_termination:
    executables:
      - context: "system"
        filename: "C:\\AppStream\\Scripts\\End-System.ps1"
        arguments: "-Cleanup"
        s3LogEnabled: true
      - context: "user"
        filename: "C:\\AppStream\\Scripts\\End-User.ps1"
        arguments: "-CleanupUser"
        s3LogEnabled: true
    waitingTime: 60

image:
  name: "example_image"
  display_name: "example image"
  description: "example image"
  enable_dynamic_app_catalog: true
  use_latest_agent_version: false
  tags:
    - team:infra
    - env:dev
  dry_run: false
```
