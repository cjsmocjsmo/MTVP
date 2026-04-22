
import unittest
from unittest.mock import patch, MagicMock
import utils

import tempfile
import shutil
import os


class TestUtils(unittest.TestCase):
    def test_mtv_walk_dirs(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            files = ["a.mp4", "b.mkv", "c.avi", "d.mpg", "e.txt"]
            for fname in files:
                open(os.path.join(tmpdir, fname), "w").close()
            result = utils.mtv_walk_dirs(tmpdir)
            self.assertCountEqual(
                [os.path.join(tmpdir, f) for f in files if f.endswith((".mp4", ".mkv", ".avi", ".mpg"))],
                result
            )

    def test_img_walk_dirs(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            files = ["a.jpg", "b.JPG", "c.png", "d.txt"]
            for fname in files:
                open(os.path.join(tmpdir, fname), "w").close()
            result = utils.img_walk_dirs(tmpdir)
            self.assertEqual([os.path.join(tmpdir, "a.jpg")], result)

    def test_tv_img_walk_dirs(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            files = ["a.jpg", "b.jpg", "c.txt"]
            for fname in files:
                open(os.path.join(tmpdir, fname), "w").close()
            result = utils.tv_img_walk_dirs(tmpdir)
            self.assertCountEqual([os.path.join(tmpdir, "a.jpg"), os.path.join(tmpdir, "b.jpg")], result)

    @patch("utils.subprocess.run")
    def test_sqlite3_check(self, mock_run):
        mock_run.return_value.returncode = 0
        self.assertTrue(utils.sqlite3_check())
        mock_run.return_value.returncode = 1
        self.assertFalse(utils.sqlite3_check())

    @patch("utils.subprocess.run")
    def test_mpv_check(self, mock_run):
        mock_run.return_value.returncode = 0
        self.assertTrue(utils.mpv_check())
        mock_run.return_value.returncode = 1
        self.assertFalse(utils.mpv_check())

    @patch("utils.subprocess.run")
    def test_python3_pil_check(self, mock_run):
        mock_run.return_value.returncode = 0
        self.assertTrue(utils.python3_pil_check())
        mock_run.return_value.returncode = 1
        self.assertFalse(utils.python3_pil_check())

    @patch("utils.subprocess.run")
    def test_python3_dotenv_check(self, mock_run):
        mock_run.return_value.returncode = 0
        self.assertTrue(utils.python3_dotenv_check())
        mock_run.return_value.returncode = 1
        self.assertFalse(utils.python3_dotenv_check())

    @patch("utils.subprocess.run")
    def test_python3_websockets_check(self, mock_run):
        mock_run.return_value.returncode = 0
        self.assertTrue(utils.python3_websockets_check())
        mock_run.return_value.returncode = 1
        self.assertFalse(utils.python3_websockets_check())

    @patch("builtins.__import__")
    def test_python3_mpv_check(self, mock_import):
        # Simulate import mpv success
        def import_side_effect(name, *args, **kwargs):
            if name == "mpv":
                return None
            return __import__(name, *args, **kwargs)
        mock_import.side_effect = import_side_effect
        self.assertTrue(utils.python3_mpv_check())
        # Simulate ImportError
        mock_import.side_effect = ImportError
        self.assertFalse(utils.python3_mpv_check())

    def test_get_arch(self):
        arch = utils.get_arch()
        self.assertIn(arch, ["32", "64", None])

    @patch("utils.sqlite3.connect")
    def test_movie_count(self, mock_connect):
        mock_conn = MagicMock()
        mock_cursor = MagicMock()
        mock_connect.return_value = mock_conn
        mock_conn.cursor.return_value = mock_cursor
        mock_cursor.fetchone.return_value = [42]
        count = utils.movie_count()
        self.assertEqual(count, 42)

    @patch("utils.sqlite3.connect")
    def test_tvshow_count(self, mock_connect):
        mock_conn = MagicMock()
        mock_cursor = MagicMock()
        mock_connect.return_value = mock_conn
        mock_conn.cursor.return_value = mock_cursor
        mock_cursor.fetchone.return_value = [24]
        count = utils.tvshow_count()
        self.assertEqual(count, 24)

    @patch("utils.sqlite3.connect")
    def test_movies_size_on_disk(self, mock_connect):
        mock_conn = MagicMock()
        mock_cursor = MagicMock()
        mock_connect.return_value = mock_conn
        mock_conn.cursor.return_value = mock_cursor
        mock_cursor.fetchall.return_value = [(1073741824,), (2147483648,)]
        size = utils.movies_size_on_disk()
        self.assertAlmostEqual(size, 3.0, places=1)

    @patch("utils.sqlite3.connect")
    def test_tvshows_size_on_disk(self, mock_connect):
        mock_conn = MagicMock()
        mock_cursor = MagicMock()
        mock_connect.return_value = mock_conn
        mock_conn.cursor.return_value = mock_cursor
        mock_cursor.fetchall.return_value = [(536870912,), (536870912,)]
        size = utils.tvshows_size_on_disk()
        self.assertAlmostEqual(size, 1.0, places=1)

if __name__ == "__main__":
    unittest.main()
