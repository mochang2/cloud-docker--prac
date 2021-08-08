from flask import Flask

app = Flask(__name__)

@app.route("/")
def home():
    return "Hello"

# FLASK_DEBUG=true FLASK_APP=main flask run => debug mode
# or app.run(debug=True, ~~)
# but only 127.0.0.1 can be used -> port-forwarding needed
if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)