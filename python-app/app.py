import yaml
import requests
from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello():
    return "Hello, World! (Python)"


@app.route("/config")
def config():
    sample_yaml = "greeting: hello\n"
    return yaml.safe_load(sample_yaml)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
