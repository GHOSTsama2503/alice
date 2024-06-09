import json
import logging
import os

from .. import env


log = logging.getLogger(__name__)

locales: list[str] = []
strings: dict[str, str] = {}


def load_locales():
    for language_code in os.listdir(env.LOCALES_DIR):
        locale_path = os.path.join(env.LOCALES_DIR, language_code)

        locales.append(language_code)
        strings[language_code] = {}

        for locale_file in os.listdir(locale_path):
            file_path = os.path.join(locale_path, locale_file)

            with open(file_path, "rb") as file:
                module_name = locale_file.split(".")[0]
                strings[language_code][module_name] = json.loads(file.read())


def user_language(user_id: int) -> str:
    return "es"


def get_locale(user_id: int, key: str) -> str:
    try:
        lang = user_language(user_id)
        data = strings.get(lang)

        for i in key.split("."):
            data = data.get(i)
        return data

    except:
        message = f"Language Key Not Found: {key}"
        log.warning(message)
        raise LanguageKeyError(message)


def bot_commands(language_code: str) -> dict[str, str]:
    return strings[language_code]["commands"]


class LanguageKeyError(Exception):
    def __init__(self, message: str) -> None:
        self.message = message

    def __str__(self) -> str:
        return self.message
