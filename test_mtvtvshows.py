

import unittest
from unittest.mock import patch, MagicMock

with patch("logging.basicConfig"), patch("logging.FileHandler"):
    import mtvtvshows

class TestProcessTVShows(unittest.TestCase):
    def setUp(self):
        self.mock_conn = MagicMock()
        self.mock_cursor = MagicMock()
        self.tvlist = ["/tmp/Show S01E01.mkv"]
        self.ptv = mtvtvshows.ProcessTVShows(self.tvlist, self.mock_conn, self.mock_cursor)

    def test_class_init(self):
        self.assertIsInstance(self.ptv, mtvtvshows.ProcessTVShows)

    def test_get_tvid(self):
        result = self.ptv.get_tvid("Show S01E01.mkv")
        self.assertIsInstance(result, str)

    def test_get_catagory(self):
        result = self.ptv.get_catagory("Shogun S01E01.mkv")
        self.assertEqual(result, "Shogun")

    def test_get_name(self):
        with patch("mtvtvshows.os.path.split", return_value=("/tmp", "Show S01E01.mkv")):
            result = self.ptv.get_name("/tmp/Show S01E01.mkv")
            self.assertIsInstance(result, str)

    def test_get_season(self):
        # Use a filename similar to real files: 'Test S01E01 Episode01.mkv'
        result = self.ptv.get_season("Test S01E01 Episode01.mkv")
        self.assertEqual(result, "01")

    def test_get_episode(self):
        # Use a filename similar to real files: 'Test S01E01 Episode01.mkv'
        result = self.ptv.get_episode("Test S01E01 Episode01.mkv")
        self.assertEqual(result, "01")

    @patch("mtvtvshows.os.stat")
    def test_process(self, mock_stat):
        mock_stat.return_value.st_size = 123
        self.ptv.cursor.execute = MagicMock()
        self.ptv.conn.commit = MagicMock()
        self.ptv.process()
        self.assertTrue(self.ptv.cursor.execute.called)

class TestUpdateTVShows(unittest.TestCase):
    def setUp(self):
        self.mock_conn = MagicMock()
        self.mock_cursor = MagicMock()
        self.upd = mtvtvshows.UpdateTVShows(self.mock_conn, self.mock_cursor)

    def test_class_init(self):
        self.assertIsInstance(self.upd, mtvtvshows.UpdateTVShows)

    def test_tvshow_paths_from_db(self):
        self.upd.cursor.fetchall.return_value = [("/tmp/Show1.mkv",), ("/tmp/Show2.mkv",)]
        result = self.upd.tvshow_paths_from_db()
        self.assertEqual(result, ["/tmp/Show1.mkv", "/tmp/Show2.mkv"])

    @patch("mtvtvshows.os.walk", return_value=[("/tmp", [], ["Show1.mkv", "Show2.mkv"])])
    @patch("mtvtvshows.os.getenv", return_value="/tmp")
    def test_tvshow_paths_from_disk(self, mock_getenv, mock_walk):
        result = self.upd.tvshow_paths_from_disk()
        self.assertIn("/tmp/Show1.mkv", result)

    @patch.object(mtvtvshows.UpdateTVShows, "tvshow_paths_from_db", return_value=["/tmp/Show1.mkv"])
    @patch.object(mtvtvshows.UpdateTVShows, "tvshow_paths_from_disk", return_value=["/tmp/Show2.mkv"])
    def test_check_for_tv_updates(self, mock_disk, mock_db):
        result = self.upd.check_for_tv_updates()
        self.assertEqual(result, ["/tmp/Show2.mkv"])

    @patch.object(mtvtvshows.UpdateTVShows, "check_for_tv_updates", return_value=["/tmp/Show2.mkv"])
    @patch("mtvtvshows.logging.info")
    @patch("mtvtvshows.ProcessTVShows.process")
    def test_updateTV(self, mock_proc, mock_log, mock_check):
        self.upd.updateTV()
        mock_log.assert_called()

if __name__ == "__main__":
    unittest.main()
