#
# Copyright (C) 2024 by TheTeamVivek@Github, < https://github.com/TheTeamVivek >.
#
# This file is part of < https://github.com/TheTeamVivek/YukkiMusic > project,
# and is released under the MIT License.
# Please see < https://github.com/TheTeamVivek/YukkiMusic/blob/master/LICENSE >
#
# All rights reserved.
#
from config import BANNED_USERS
from strings import get_command
from YukkiMusic import app
from YukkiMusic.core.call import Yukki
from YukkiMusic.utils.database import is_music_playing, music_on
from YukkiMusic.utils.decorators import AdminRightsCheck

RESUME_COMMAND = get_command("RESUME_COMMAND")


@app.on_message(
    command=RESUME_COMMAND,
    is_group=True,
    from_user=BANNED_USERS,
    is_restricted=True,
)
@AdminRightsCheck
async def resume_com(event, _, chat_id):
    if len(event.message.text.split()) != 1:
        await event.reply(_["general_2"])
        return
    if await is_music_playing(chat_id):
        await event.reply(_["admin_3"])
        return
    await music_on(chat_id)
    await Yukki.resume_stream(chat_id)

    sender = await event.get_sender()
    mention = f"[{sender.first_name}](tg://user?id={sender.id})"
    await event.reply(_["admin_4"].format(mention))
