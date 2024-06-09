#!/bin/python3
from pyrogram.client import Client

def generate_string_session(api_id: int | str, api_hash: str, bot_token: str) -> None:
    with Client("bot", api_id, api_hash, in_memory = True, bot_token = bot_token) as client:
        return client.export_session_string()

if __name__ == "__main__":
    API_ID = input("API_ID: ")
    API_HASH = input("API_HASH: ")
    BOT_TOKEN = input("BOT_TOKEN: ")

    print("\n", generate_string_session(API_ID, API_HASH, BOT_TOKEN))