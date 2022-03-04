# export DOCKER_HOST=ssh://ubuntu@192.168.1.98

bash ./_devtools/subtask/01-run-base-service.sh

echo 'We need to wait for 20s before running test case...'
sleep 20

# Build product image and test
bash ./_devtools/subtask/02-build-todo-svc.sh

# run test
bash ./_devtools/subtask/03-run-test.sh
