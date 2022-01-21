from flask import Flask, request, render_template

from database import init_databases

print("database loading...")
init_databases()
print("database loaded")

app = Flask(__name__)


@app.route("/")
def main():
    st = "Hello test"
    return st


if __name__ == '__main__':
    app.run(debug=True)
