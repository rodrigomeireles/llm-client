# TODO implementar rendering bonito de Markdown
# TODO parar de dar overflow na caixinha
import markdown
from fastapi import FastAPI, HTTPException, Form, Request

# import openai
from groq import Groq
from fastapi.staticfiles import StaticFiles
from typing import Annotated
from starlette.middleware.sessions import SessionMiddleware
from fastapi.responses import HTMLResponse

app = FastAPI()
FRONT_HTML = "static/index.html"
client = Groq()


def list_to_li(messages: list) -> str:
    response = ""
    html_messages = [
        markdown.markdown(message, extensions=["fenced_code"]) for message in messages
    ]
    for idx, message in enumerate(html_messages):
        message = message.replace("<p>", "", 1).replace("</p>", "", 1)
        if idx % 2 == 0:
            response += "<li> Eu: " + message + "</li>"
            print(message)
        else:
            response += "<li> Chat: " + message + "</li>"
    return response


def list_to_gpt_list(messages: list, top_k: int = 10) -> list[dict[str, str]]:
    gpt_messages = []
    for idx, message in enumerate(messages):
        if idx > top_k - 1:
            break
        if idx % 2 == 0:
            gpt_messages.append({"role": "user", "content": message})
        else:
            gpt_messages.append({"role": "assistant", "content": message})
    return gpt_messages


@app.get("/")
async def read_root():
    return {"Hello": "World"}


def generate_response(history: list[str]):
    if not history:
        raise HTTPException(status_code=400, detail="Message cannot be empty")
    formatted_history = list_to_gpt_list(messages=history)
    response = client.chat.completions.create(
        model="mixtral-8x7b-32768",
        messages=formatted_history,
        temperature=0.5,
        top_p=0,
        # max_tokens=2500,  # You can adjust this as needed
    )
    print(formatted_history)
    if response.choices:  # type: ignore
        generated_text = response.choices[0].message  # type: ignore
        return {"chat_response": generated_text.content}
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
        messages = []
    messages.append(user_input)
    response = generate_response(history=messages)
    chat_response = response.get("chat_response")
    if chat_response:
        messages.append(chat_response)
    sess["messages"] = messages
    return list_to_li(sess["messages"])


@app.get("/get_history", response_class=HTMLResponse)
async def get_history(request: Request):
    sess = request.session
    messages = sess.get("messages")
    if not messages:
        sess["messages"] = []
    return list_to_li(sess["messages"])


@app.put("/refresh_session")
async def refresh(request: Request, response_class=HTMLResponse):
    request.session.clear()
    return response_class("<p>Sessão limpa.</p>")


app.mount("/", StaticFiles(directory="static"), name="static")
app.add_middleware(SessionMiddleware, secret_key="key")
