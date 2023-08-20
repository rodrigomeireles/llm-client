from fastapi import FastAPI, HTTPException
import openai

app = FastAPI()


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

    if response.choices:
        generated_text = response.choices[0].message
        return {"response": generated_text}
    else:
        raise HTTPException(status_code=500, detail="Failed to generate response")
