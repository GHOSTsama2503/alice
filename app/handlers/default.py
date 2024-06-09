from pyrogram import Client
from pyrogram.enums import ParseMode
from pyrogram.types import InlineKeyboardButton, InlineKeyboardMarkup, Message

from .. import env
from ..modules import language
from ..utils import filters


@Client.on_message(filters.command("start"))
async def start_command_handler(client: Client, message: Message):
    """Welcome message"""

    welcome_message = language.get_locale(
        message.from_user.id, "default.welcome_message"
    ).format(
        name=message.from_user.mention(style=ParseMode.MARKDOWN),
        client=env.CLIENT_NAME.title(),
        image=f"[ㅤ]({env.BOT_PICTURE})" if env.BOT_PICTURE else "",
    )

    change_language = language.get_locale(
        message.from_user.id, "default.change_language"
    )

    language_button = [InlineKeyboardButton(change_language, "*")]
    reply_markup = InlineKeyboardMarkup([language_button])

    await message.reply(
        welcome_message,
        quote=True,
        reply_markup=reply_markup,
        parse_mode=ParseMode.MARKDOWN,
    )
