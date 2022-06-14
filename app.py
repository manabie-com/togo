from service.service import app

if __name__ == '__main__':
    app.run(port=8000, host="0.0.0.0", debug=True)
