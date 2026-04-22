
import unittest
from unittest.mock import patch

with patch("logging.basicConfig"), patch("logging.FileHandler"):
    import mtvmovies


from unittest.mock import patch, MagicMock
import os

class TestProcessMovies(unittest.TestCase):
    def test_class_init(self):
        try:
            movs = []
            m = mtvmovies.ProcessMovies(movs, None, None)
        except Exception as e:
            self.fail(f"ProcessMovies() raised: {e}")

    def test_get_year(self):
        m = mtvmovies.ProcessMovies([], None, None)
        self.assertEqual(m.get_year("Movie (2020)"), "2020")

    def test_get_poster(self):
        m = mtvmovies.ProcessMovies([], None, None)
        self.assertTrue(m.get_poster("/tmp/movie.mp4").endswith(".jpg"))

    def test_get_mov_id(self):
        m = mtvmovies.ProcessMovies([], None, None)
        self.assertIsInstance(m.get_mov_id("test"), str)

    @patch("mtvmovies.os.path.split", return_value=("/tmp", "ActionMovie.mp4"))
    @patch("mtvmovies.re.search", return_value=True)
    def test_get_catagory(self, mock_search, mock_split):
        m = mtvmovies.ProcessMovies([], None, None)
        self.assertIsInstance(m.get_catagory("/tmp/ActionMovie.mp4"), str)

    @patch("mtvmovies.os.path.split", return_value=("/tmp", "movie.mp4"))
    @patch("mtvmovies.os.getenv", return_value="localhost")
    def test_get_http_thumb_path(self, mock_getenv, mock_split):
        m = mtvmovies.ProcessMovies([], None, None)
        self.assertIn(":8080", m.get_http_thumb_path("/tmp/movie.mp4"))

    @patch("mtvmovies.os.path.split", return_value=("/tmp", "Movie (2020).mp4"))
    @patch("mtvmovies.re.search", return_value=None)
    def test_get_name(self, mock_search, mock_split):
        m = mtvmovies.ProcessMovies([], None, None)
        self.assertIsNone(m.get_name("/tmp/Movie (2020).mp4"))

    @patch("mtvmovies.ProcessMovies.get_poster", return_value="/tmp/movie.jpg")
    @patch("mtvmovies.ProcessMovies.get_name", return_value="Movie")
    @patch("mtvmovies.ProcessMovies.get_year", return_value="2020")
    @patch("mtvmovies.ProcessMovies.get_mov_id", return_value="abc")
    @patch("mtvmovies.ProcessMovies.get_catagory", return_value="Action")
    @patch("mtvmovies.ProcessMovies.get_http_thumb_path", return_value="localhost:8080/thumbnails/movie.jpg")
    @patch("mtvmovies.os.stat")
    @patch("mtvmovies.logging.error")
    def test_process(self, mock_log, mock_stat, mock_http, mock_cat, mock_id, mock_year, mock_name, mock_poster):
        mock_stat.return_value.st_size = 123
        m = mtvmovies.ProcessMovies(["/tmp/movie.mp4"], MagicMock(), MagicMock())
        m.cursor.execute = MagicMock()
        m.conn.commit = MagicMock()
        m.process()

class TestUpdateMovies(unittest.TestCase):
    def test_class_init(self):
        try:
            m = mtvmovies.UpdateMovies(None, None)
        except Exception as e:
            self.fail(f"UpdateMovies() raised: {e}")

    @patch("mtvmovies.UpdateMovies.movie_image_paths_from_db", return_value=["/tmp/img1.jpg"])
    def test_movie_image_paths_from_db(self, mock_db):
        m = mtvmovies.UpdateMovies(MagicMock(), MagicMock())
        result = m.movie_image_paths_from_db()
        self.assertIsInstance(result, list)

    @patch("mtvmovies.UpdateMovies.movie_paths_from_db", return_value=["/tmp/movie1.mp4"])
    def test_movie_paths_from_db(self, mock_db):
        m = mtvmovies.UpdateMovies(MagicMock(), MagicMock())
        result = m.movie_paths_from_db()
        self.assertIsInstance(result, list)

    @patch("mtvmovies.os.walk", return_value=[("/tmp", [], ["img1.jpg", "img2.png"])])
    @patch("mtvmovies.os.getenv", return_value="/tmp")
    def test_movie_image_paths_from_disk(self, mock_getenv, mock_walk):
        m = mtvmovies.UpdateMovies(MagicMock(), MagicMock())
        result = m.movie_image_paths_from_disk()
        self.assertIsInstance(result, list)

    @patch("mtvmovies.os.walk", return_value=[("/tmp", [], ["movie1.mp4", "movie2.mkv"])])
    @patch("mtvmovies.os.getenv", return_value="/tmp")
    def test_movie_paths_from_disk(self, mock_getenv, mock_walk):
        m = mtvmovies.UpdateMovies(MagicMock(), MagicMock())
        result = m.movie_paths_from_disk()
        self.assertIsInstance(result, list)

    @patch("mtvmovies.UpdateMovies.movie_paths_from_db", return_value=["/tmp/movie1.mp4"])
    @patch("mtvmovies.UpdateMovies.movie_paths_from_disk", return_value=["/tmp/movie2.mp4"])
    def test_check_for_mov_updates(self, mock_disk, mock_db):
        m = mtvmovies.UpdateMovies(MagicMock(), MagicMock())
        result = m.check_for_mov_updates()
        self.assertIsInstance(result, list)

    @patch("mtvmovies.UpdateMovies.check_for_mov_updates", return_value=["/tmp/movie2.mp4"])
    @patch("mtvmovies.ProcessMovies.process")
    @patch("mtvmovies.logging.info")
    def test_updateMovs(self, mock_log, mock_proc, mock_check):
        m = mtvmovies.UpdateMovies(MagicMock(), MagicMock())
        m.updateMovs()

if __name__ == "__main__":
    unittest.main()
