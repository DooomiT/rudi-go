.PHONY: install-stt download-models download-audio-files


install-stt:
	curl -O https://github.com/coqui-ai/STT/releases/download/v1.3.0/native_client.tflite.macOS.tar.xz
	rm -rf ${HOME}/.coqui
	mkdir ${HOME}/.coqui
	tar -xvzf native_client.tflite.macOS.tar.xz --directory ${HOME}/.coqui/
	rm native_client.tflite.macOS.tar.xz

download-models:
	if [ -d "model" ]; then echo "models already exist" && exit 0; fi
	mkdir -p model
	curl -o model/model.tflite "https://objects.githubusercontent.com/github-production-release-asset-2e65be/351871871/7f2559a8-7a9c-43d8-8d6c-233d4cf50da0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20221210%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20221210T000903Z&X-Amz-Expires=300&X-Amz-Signature=c78e660ed136b18ef516115d94e2899ee5bfab802746c51ffa9a5dd4486da5f5&X-Amz-SignedHeaders=host&actor_id=51324632&key_id=0&repo_id=351871871&response-content-disposition=attachment%3B%20filename%3Dmodel.tflite&response-content-type=application%2Foctet-stream"
	curl -o model/huge-vocabulary.scorer "https://objects.githubusercontent.com/github-production-release-asset-2e65be/351871871/8a777b8e-ad5f-4fc0-8dc3-23b0792446c0?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20221210%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20221210T001319Z&X-Amz-Expires=300&X-Amz-Signature=b3df348fb4edd21b1b2768e485b9e1708f7b4d36dfc77830ea52b17a8864a7c1&X-Amz-SignedHeaders=host&actor_id=51324632&key_id=0&repo_id=351871871&response-content-disposition=attachment%3B%20filename%3Dhuge-vocabulary.scorer&response-content-type=application%2Foctet-stream"
	
download-audio-files:
	if [ -d "audio-files" ]; then echo "models already exist" && exit 0; fi
	mkdir -p audio-files
	curl -o audio-files/audio-1.4.0.tar.gz "https://github.com/coqui-ai/STT/releases/download/v1.4.0/audio-1.4.0.tar.gz"
	tar -xvzf audio-files/audio-1.4.0.tar.gz --directory audio-files/