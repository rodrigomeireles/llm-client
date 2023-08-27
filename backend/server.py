from fastapi import FastAPI, HTTPException, Form, Request
import openai
from fastapi.staticfiles import StaticFiles
from typing import Annotated
from starlette.middleware.sessions import SessionMiddleware
from fastapi.responses import HTMLResponse

app = FastAPI()
FRONT_HTML = "static/index.html"


@app.get("/")
async def read_root():
    return {"Hello": "World"}


@app.post("/chat")
async def generate_response(data: dict):
    message = data.get("message")
    if not message:
        raise HTTPException(status_code=400, detail="Message cannot be empty")

    response = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[{"role": "user", "content": message}],
        max_tokens=50,  # You can adjust this as needed
    )

    if response.choices:  # type: ignore
        generated_text = response.choices[0].message  # type: ignore
        return {"response": generated_text}
    else:
        raise HTTPException(status_code=500, detail="Failed to generate response")


@app.get("/clear_input", response_class=HTMLResponse)
async def return_nothing():
    return HTMLResponse("")


@app.post("/send_message", response_class=HTMLResponse)
async def concatenate_message(request: Request, user_input: Annotated[str, Form()]):
    sess = request.session
    messages = sess.get("messages")
    if not messages:
        messages = ""
    messages += "<li>" + user_input + "</li>"
    sess["messages"] = messages
    return sess["messages"]


@app.get("/get_history", response_class=HTMLResponse)
async def get_history(request: Request):
    sess = request.session
    messages = sess.get("messages")
    if not messages:
        sess["messages"] = ""
    return sess["messages"]


@app.put("/refresh_session")
async def refresh(request: Request, response_class=HTMLResponse):
    request.session.clear()
    return response_class("<p>Sessão limpa.</p>")


app.mount("/", StaticFiles(directory="static"), name="static")
app.add_middleware(SessionMiddleware, secret_key="key")
