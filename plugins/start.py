from pyrogram import Client, filters
from pyrogram.types import Message, InlineKeyboardMarkup, InlineKeyboardButton
from pyrogram.handlers import MessageHandler
from pyrogram.enums.parse_mode import ParseMode


async def start_command(client: Client, event: Message, t=MessageHandler, f=filters.command("start")):
    if not event.from_user:
        return
    msg = f"Hola [{event.from_user.first_name}](tg://user?id={event.from_user.id}), "
    msg += "soy Alice, una asistente multifuncional, ¿qué puedo hacer por usted?"
    msg += "[ㅤ](https://telegra.ph/file/fdc88663a5eb402f46d84.jpg)"
    await event.reply(
        text = msg,
        reply_markup = InlineKeyboardMarkup([
            [InlineKeyboardButton("♥️ Alice desu >~<", callback_data="start")]
        ]),
        parse_mode = ParseMode.MARKDOWN,
        reply_to_message_id = event.id
    )
