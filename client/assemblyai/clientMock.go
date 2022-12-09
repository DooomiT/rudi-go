package assemblyai

type AssemblyAiMock struct {
	UploadLocalFileMock func() (string, error)
	TranscriptMock      func() (string, error)
	PollTranscriptMock  func() (string, error)
}

func (client *AssemblyAiMock) UploadLocalFile(content []byte) (string, error) {
	return client.UploadLocalFileMock()
}

func (client *AssemblyAiMock) Transcript(audioUrl string) (string, error) {
	return client.TranscriptMock()
}

func (client *AssemblyAiMock) PollTranscript(id string, maxTries uint) (string, error) {
	return client.PollTranscriptMock()
}
func mockFunction(data string, err error) func() (string, error) {
	return func() (string, error) {
		return data, err
	}
}

func NewMock(uploadFileUrl string, uploadFileError error, transcribedText string, transcribedTextError error, pollText string, pollError error) AssemblyAi {
	return &AssemblyAiMock{
		UploadLocalFileMock: mockFunction(uploadFileUrl, uploadFileError),
		TranscriptMock:      mockFunction(transcribedText, transcribedTextError),
		PollTranscriptMock:  mockFunction(pollText, pollError),
	}
}
