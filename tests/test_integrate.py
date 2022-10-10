
# test creation within limit of the user up to the deletion
def test_creation_deletion(client):
    # login and retrieve the token
    response = client.post("/login", auth=("sample_user", "sample_password"))
    token = response.get_json()["token"]

    # create a sample task
    sample_test = {"todo": "sample todo task"}
    response = client.post("/todo", json=sample_test, headers={"x-access-token": token})

    # task should be added successfully
    assert response.status_code == 200
    assert response.get_json()["message"] == "new task added."

    # exceed allocated task per day
    sample_test = {"todo": "sample todo task"}
    response = client.post("/todo", json=sample_test, headers={"x-access-token": token})

    # error code should be thrown since limit has been reached
    assert response.status_code == 400
    assert response.text == "User has reached maximum todos per day."

    # retrieve the todo

    response = client.get("/todos", headers={"x-access-token": token})

    todos = response.get_json()["todos"]

    print(todos)

    # check if the data of the first record
    assert todos[0]["task"] == "sample todo task"

    # test the deletion of a task

    response = client.get("/todos", headers={"x-access-token": token})

    todos = response.get_json()["todos"]

    # retrieve the id of the first todo item
    todo_id = todos[0]["id"]

    response = client.post("/todo/delete", json={"todo_id": todo_id},
                             headers={"x-access-token": token})

    # item should be deleted successfully
    assert response.status_code == 200
    assert response.get_json()["message"] == "todo deleted."

    # test the deletion of a user
    response = client.post("/user/delete", headers={"x-access-token": token})

    # user should be deleted successfully
    assert response.status_code == 200
    assert response.get_json()["message"] == "user deleted."
