from flask import Flask, Response

app = Flask(__name__)
MESSAGE = 'hello world!'


@app.route('/')
def hello():
    return Response(MESSAGE)


if __name__ == '__main__':
    app.run(threaded=True, port=5011)

