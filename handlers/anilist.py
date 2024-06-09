import lib.env as env
from pyrogram import Client, filters
from pyrogram.types import Message


@Client.on_message(~ filters.channel & ~ filters.group & filters.command("anime", prefixes=env.PREFIXES))
async def anime_command(client: Client, message: Message):
    pass


@Client.on_message(~ filters.channel & ~ filters.group & filters.command("manga", prefixes=env.PREFIXES))
async def manga_command(client: Client, message: Message):
    pass


@Client.on_message(~ filters.channel & ~ filters.group & filters.command("airing", prefixes=env.PREFIXES))
async def airing_command(client: Client, message: Message):
    pass

