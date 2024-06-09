import subprocess

from pyrogram import Client
from pyrogram.enums import ParseMode
from pyrogram.types import Message

from ..utils import filters


@Client.on_message(filters.command("sh") & filters.only_admins())
async def sh_shell_command(client: Client, message: Message):
    """Execute commands in the terminal and return the output of the same"""

    # Extract command from user input
    if message.text[1:4] == "sh ":
        command: str = message.text[4:]
    elif (
        len(message.text) == 3
        and message.reply_to_message
        and message.reply_to_message.text
    ):
        command: str = message.reply_to_message.text
    else:
        return

    # Run the command
    process = subprocess.Popen(
        command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, shell=True
    )
    stdout, stderr = process.communicate()
    stdout = stdout.decode("UTF-8")
    stderr = stderr.decode("UTF-8")

    reply_text = f"**Input:**\n`{command}`\n\n"
    if stdout:
        reply_text += f"**Output:**\n`{stdout}`\n\n"
    if stderr:
        reply_text += f"**Error:**\n`{stderr}`"

    output_length_limit = 4000

    # If the output exceeds the limit insert it into a file
    if len(reply_text) >= output_length_limit:
        output_file = "./stdout.txt"
        with open(output_file, "w") as file:
            file.write(stdout)

        caption = f"**Input:**\n`{command}`"
        await message.reply_document(
            output_file, caption=caption, quote=True, parse_mode=ParseMode.MARKDOWN
        )

    # If not, reply with a text message
    else:
        await message.reply(reply_text, quote=True, parse_mode=ParseMode.MARKDOWN)
