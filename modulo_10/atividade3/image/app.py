from rembg import remove
from PIL import Image
from fastapi import FastAPI, File, UploadFile
import io

app = FastAPI()

@app.post("/remove-background/")
async def remove_background(file: UploadFile = File(...)):
    input_image = Image.open(io.BytesIO(await file.read()))
    output_image = remove(input_image)
    
    output_buffer = io.BytesIO()
    output_image.save(output_buffer, format="PNG")
    output_buffer.seek(0)
    
    return {
        "filename": file.filename,
        "content": output_buffer.read()
    }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
