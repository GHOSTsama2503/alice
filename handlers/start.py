import lib.env as env
from modules.language import strings
from pyrogram import Client, filters
from pyrogram.enums import ParseMode
from pyrogram.types import InlineKeyboardButton, InlineKeyboardMarkup, Message


@Client.on_message(~ filters.channel & ~ filters.group & filters.command("start", prefixes=env.PREFIXES))
async def starts_command(client: Client, message: Message):
    """Respond with a welcome message to the start's command"""

    name = f"[{message.from_user.first_name}](tg://user?id={message.from_user.id})"
    image = "[ㅤ](https://telegra.ph/file/fdc88663a5eb402f46d84.jpg)"
    lang = message.from_user.language_code

    button = [InlineKeyboardButton(strings[lang]["welcome_message_button"], ".")]
    reply_markup = InlineKeyboardMarkup([button])

    text = strings[lang]["welcome_message"].format(name=name, image=image)

    await message.reply(text, quote=True, reply_markup=reply_markup, parse_mode=ParseMode.MARKDOWN)
