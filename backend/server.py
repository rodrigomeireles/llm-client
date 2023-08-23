from fastapi import FastAPI, HTTPException, Form
import openai
from fastapi.staticfiles import StaticFiles
from typing import Annotated

app = FastAPI()
FRONT_HTML = "static/index.html"
messages = ""


@app.get("/")
def read_root():
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


# TODO acho que o que eu quero é um WebSocket (API Rest não tem estado).
# TODO adicionar um endpoint WebSocket que recebe um Cookie da sessão.
@app.post("/send_message")
async def concatenate_message(user_input: Annotated[str, Form()]):
    messages += "<p>" + user_input + "</p>"
    print(messages)
    return messages


app.mount("/", StaticFiles(directory="static"), name="static")
