from transformers import ElectraForSequenceClassification, ElectraTokenizer
import torch

# ELECTRA 모델 및 토크나이저 초기화
model_name = "google/electra-small-discriminator"
model = ElectraForSequenceClassification.from_pretrained(model_name)
tokenizer = ElectraTokenizer.from_pretrained(model_name)

# 입력 문장
input_text = "This is an example sentence."

# 문장 토큰화
input_ids = tokenizer.encode(input_text,return_tensors="pt")

outputs = model(input_ids)

predictions = outputs.logits

print(predictions)