import app

if __name__ == "__main__":
    main_app = app.create_app()
    app.init_routes(main_app)
    main_app.run(debug=True)
