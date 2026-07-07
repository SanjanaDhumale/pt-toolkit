[Setup]
AppId={{PTToolkit}}
AppName=PT Toolkit
AppVersion=1.0
AppPublisher=Sanjana Dhumale
DefaultDirName={autopf}\PT Toolkit
DefaultGroupName=PT Toolkit

ArchitecturesAllowed=x64compatible
ArchitecturesInstallIn64BitMode=x64compatible

OutputDir=..\releases\v1.0
OutputBaseFilename=PTToolkitSetup
Compression=lzma
SolidCompression=yes
WizardStyle=modern

[Tasks]
Name: "desktopicon"; Description: "Create a Desktop Shortcut"; GroupDescription: "Additional Icons:"

[Files]

; Main executable
Source: "..\releases\v1.0\ptctl.exe"; DestDir: "{app}"; Flags: ignoreversion

; Documentation
Source: "..\releases\v1.0\README.md"; DestDir: "{app}"; Flags: ignoreversion

; Docker files
Source: "..\releases\v1.0\docker\*"; DestDir: "{app}\docker"; Flags: ignoreversion recursesubdirs createallsubdirs

; Documentation
Source: "..\releases\v1.0\docs\*"; DestDir: "{app}\docs"; Flags: ignoreversion recursesubdirs createallsubdirs

[Icons]
Name: "{group}\PT Toolkit"; Filename: "{app}\ptctl.exe"
Name: "{autodesktop}\PT Toolkit"; Filename: "{app}\ptctl.exe"; Tasks: desktopicon