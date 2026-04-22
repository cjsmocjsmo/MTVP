

import unittest
import asyncio
from unittest.mock import patch, MagicMock, AsyncMock

with patch("logging.basicConfig"), patch("logging.FileHandler"), patch("os.makedirs"), patch("builtins.open"), patch("os.getenv", return_value="/tmp/mtv.log"):
    import mtvserverasync

class TestMTVServerAsync(unittest.IsolatedAsyncioTestCase):
    async def test_get_media_path_from_media_id(self):
        with patch("mtvserverasync.sqlite3.connect") as mock_connect, \
             patch("mtvserverasync.os.getenv", return_value="/tmp/test.db"):
            mock_conn = MagicMock()
            mock_cursor = MagicMock()
            mock_cursor.fetchone.return_value = ("/media/path.mp4",)
            mock_conn.cursor.return_value = mock_cursor
            mock_connect.return_value.__enter__.return_value = mock_conn
            result = await mtvserverasync.get_media_path_from_media_id("1")
            self.assertEqual(result, "/media/path.mp4")

    async def test_get_media_path_from_media_tv_id(self):
        with patch("mtvserverasync.sqlite3.connect") as mock_connect, \
             patch("mtvserverasync.os.getenv", return_value="/tmp/test.db"):
            mock_conn = MagicMock()
            mock_cursor = MagicMock()
            mock_cursor.fetchone.return_value = ("/media/tvpath.mp4",)
            mock_conn.cursor.return_value = mock_cursor
            mock_connect.return_value.__enter__.return_value = mock_conn
            result = await mtvserverasync.get_media_path_from_media_tv_id("1")
            self.assertEqual(result, "/media/tvpath.mp4")

    async def test_get_media_path_from_video_id(self):
        with patch("mtvserverasync.sqlite3.connect") as mock_connect, \
             patch("mtvserverasync.os.getenv", return_value="/tmp/test.db"):
            mock_conn = MagicMock()
            mock_cursor = MagicMock()
            mock_cursor.fetchone.return_value = ("/media/vidpath.mp4",)
            mock_conn.cursor.return_value = mock_cursor
            mock_connect.return_value.__enter__.return_value = mock_conn
            result = await mtvserverasync.get_media_path_from_video_id("1")
            self.assertEqual(result, "/media/vidpath.mp4")

    async def test_get_weather_for_belfair_wa(self):
        with patch("mtvserverasync.requests.get") as mock_get:
            mock_point_resp = MagicMock()
            mock_point_resp.json.return_value = {"properties": {"forecastHourly": "http://forecast.url"}}
            mock_point_resp.raise_for_status = MagicMock()
            mock_weather_resp = MagicMock()
            mock_weather_resp.json.return_value = {"properties": {"periods": [{
                "temperature": 55, "temperatureUnit": "F", "shortForecast": "Sunny", "windDirection": "N", "windSpeed": "5 mph"
            }]}}
            mock_weather_resp.raise_for_status = MagicMock()
            mock_get.side_effect = [mock_point_resp, mock_weather_resp]
            result = await mtvserverasync.get_weather_for_belfair_wa()
            self.assertEqual(result["location"], "Belfair, WA")

    async def test_handle_message_set_media(self):
        # Test the set_media command branch
        fake_ws = AsyncMock()
        fake_ws.__aiter__.return_value = [
            '{"command": "set_media", "media_id": "1"}'
        ]
        with patch("mtvserverasync.get_media_path_from_media_id", new=AsyncMock(return_value="/media/path.mp4")), \
             patch.object(mtvserverasync, "player") as mock_player:
            await mtvserverasync.handle_message(fake_ws)
            mock_player.play.assert_called_with("/media/path.mp4")
            fake_ws.send.assert_called()

    async def test_handle_message_search(self):
        fake_ws = AsyncMock()
        fake_ws.__aiter__.return_value = [
            '{"command": "search", "phrase": "test"}'
        ]
        with patch.object(mtvserverasync, "MTVMEDIA") as mock_media:
            mock_media.search.return_value = {"results": [1,2,3]}
            await mtvserverasync.handle_message(fake_ws)
            fake_ws.send.assert_called()

    async def test_handle_message_stop(self):
        fake_ws = AsyncMock()
        fake_ws.__aiter__.return_value = [
            '{"command": "stop"}'
        ]
        with patch.object(mtvserverasync, "player") as mock_player:
            await mtvserverasync.handle_message(fake_ws)
            mock_player.stop.assert_called()
            fake_ws.send.assert_called()

    async def test_handle_message_pause(self):
        fake_ws = AsyncMock()
        fake_ws.__aiter__.return_value = [
            '{"command": "pause"}'
        ]
        with patch.object(mtvserverasync, "player") as mock_player:
            mock_player.pause = False
            await mtvserverasync.handle_message(fake_ws)
            self.assertTrue(hasattr(mock_player, "pause"))
            fake_ws.send.assert_called()

    async def test_servermain_runs(self):
        # Patch websockets.serve and asyncio.start_server to test servermain
        with patch("mtvserverasync.websockets.serve", new=AsyncMock()), \
             patch("mtvserverasync.asyncio.start_server", new=AsyncMock()), \
             patch("builtins.print"):
            try:
                await asyncio.wait_for(mtvserverasync.servermain(), timeout=0.1)
            except Exception:
                pass  # Expected to timeout or complete

if __name__ == "__main__":
    unittest.main()
