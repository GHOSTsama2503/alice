from pyrogram import filters
from .. import env
from pyrogram.types import CallbackQuery, Message


def command(commands: list[str] | str, case_sensitive: bool = False) -> filters.Filter:
    """Filter commands, i.e.: text messages starting with any custom prefix.

    Parameters:
        commands (``str`` | ``list``):
            The command or list of commands as string the filter should look for.
            Examples: "start", ["start", "help", "settings"]. When a message text containing
            a command arrives, the command itself and its arguments will be stored in the *command*
            field of the :obj:`~pyrogram.types.Message`.

        case_sensitive (``bool``, *optional*):
            Pass True if you want your command(s) to be case sensitive. Defaults to False.
            Examples: when True, command="Start" would trigger /Start but not /start.
    """
    return filters.command(
        commands=commands, prefixes=env.PREFIXES, case_sensitive=case_sensitive
    )


def from_user() -> filters.Filter:
    def _from_user(_, __, event: CallbackQuery | Message) -> bool:
        return bool(event.from_user)

    return filters.create(_from_user)


def only_admins() -> filters.Filter:
    return filters.user(env.ADMIN_ID)
