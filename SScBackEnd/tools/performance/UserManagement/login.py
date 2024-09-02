from locust import HttpUser, TaskSet, task, between
import random

class UserBehavior(TaskSet):
    emails = [
        "neerajsrivastav7@gmail.com",
        "example1@gmail.com",
        "example2@gmail.com"
    ]
    passwords = [
        "password123",
        "password456",
        "password789"
    ]
    logintypes = ["password", "email"]

    @task
    def login(self):
        email = random.choice(self.emails)
        password = random.choice(self.passwords)
        logintype = random.choice(self.logintypes)
        
        self.client.post(
            "/api/auth/login",
            json={
                "email": email,
                "password": password,
                "logintype": logintype
            },
            headers={"Content-Type": "application/json"}
        )

class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 5)  # Simulates a wait time between 1 and 5 seconds
