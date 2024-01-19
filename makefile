FIREBASE_URL = https://task-manager-fb03a-default-rtdb.asia-southeast1.firebasedatabase.app/task.json

run:
	export FIREBASE_URL=$(FIREBASE_URL) && go run main.go
