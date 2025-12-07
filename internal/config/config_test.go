package config_test

// func TestNormaliseConfig(t *testing.T) {
// 	testCases := []struct {
// 		desc        string
// 		fileContent string
// 		expected    string
// 	}{
// 		{
// 			desc: "invalid config windows",
// 			fileContent: `platform: "windows"
// files:
//   - path: "C:\\Users\\Alice\\Desktop\\..\\Temp\\.\\notes.txt"
//     content: |
//       echo "Hello World"`,
// 			expected: "C:\\Users\\Alice\\Desktop\\notes.txt",
// 		},
// 		{
// 			desc: "invalid config unix",
// 			fileContent: `platform: "unix"
// files:
//   - path: "/home/user/../user/documents/./report.pdf"
//     content: |
//       echo "Hello World"`,
// 			expected: "/home/user/documents/report.pdf",
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			file, err := os.CreateTemp("../../testdata", fmt.Sprintf("test_%s_", tC.desc))

// 			if err != nil {
// 				t.Errorf("unable to create temp file: %v", err)
// 			}

// 			file.WriteString(tC.fileContent)

// 			lb := backend.LocalBackend{
// 				Location: file.Name(),
// 			}

// 			configData, err := lb.GetConfig()

// 			if err != nil {
// 				t.Errorf("unable to fetch config data: %v", err)
// 			}

// 			// configData.NormaliseConfig()

// 			if configData.Files[0].Path != tC.expected {
// 				t.Errorf("expected %s, got %s", tC.expected, configData.Files[0].Path)
// 			}

// 			t.Cleanup(func() {
// 				os.Remove(file.Name())
// 			})

// 			t.Cleanup(func() {
// 				file.Close()
// 			})
// 		})
// 	}
// }
