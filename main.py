from fastapi import FastAPI
from fastapi.responses import RedirectResponse
from routes import router


app = FastAPI()
app.include_router(router, prefix='/tasks')


@app.get('/')
def api_docs():
    return RedirectResponse(url='/docs')
