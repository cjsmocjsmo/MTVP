
import unittest
from unittest.mock import patch

with patch("logging.basicConfig"), patch("logging.FileHandler"):
    import mtvimages


from unittest.mock import patch, MagicMock
import tempfile
import os

class TestProcessImages(unittest.TestCase):
    def test_class_init(self):
        try:
            imgs = []
            m = mtvimages.ProcessImages(imgs, None, None)
        except Exception as e:
            self.fail(f"ProcessImages() raised: {e}")

    @patch("mtvimages.os.path.exists", return_value=True)
    def test_thumb_dir_check(self, mock_exists):
        imgs = []
        m = mtvimages.ProcessImages(imgs, None, None)
        m.thumb_dir_check()

    @patch("mtvimages.Image.open")
    @patch("mtvimages.os.path.split", side_effect=lambda x: ("/tmp", "test.jpg"))
    @patch("mtvimages.os.path.join", side_effect=lambda a, b: f"{a}/{b}")
    @patch("mtvimages.os.getenv", return_value="/tmp")
    def test_create_thumbnail(self, mock_getenv, mock_join, mock_split, mock_open):
        imgs = []
        m = mtvimages.ProcessImages(imgs, None, None)
        mock_thumb = MagicMock()
        mock_open.return_value = mock_thumb
        mock_thumb.thumbnail.return_value = None
        mock_thumb.save.return_value = None
        result = m.create_thumbnail("/tmp/test.jpg")
        self.assertIn("/tmp", result)

    def test_get_img_id(self):
        imgs = []
        m = mtvimages.ProcessImages(imgs, None, None)
        result = m.get_img_id("teststring")
        self.assertIsInstance(result, str)

    @patch("mtvimages.re.search", return_value=None)
    @patch("mtvimages.os.path.split", side_effect=lambda x: ("/tmp", "test.jpg"))
    def test_get_name(self, mock_split, mock_search):
        imgs = []
        m = mtvimages.ProcessImages(imgs, None, None)
        result = m.get_name("/tmp/test.jpg")
        self.assertIsInstance(result, str)

    @patch("mtvimages.os.getenv", return_value="/tmp")
    def test_get_thumb_path(self, mock_getenv):
        imgs = []
        m = mtvimages.ProcessImages(imgs, None, None)
        result = m.get_thumb_path("/tmp/test.jpg")
        self.assertIn("/tmp", result)

    @patch("mtvimages.os.getenv", return_value="/tmp")
    def test_get_http_thumb_path(self, mock_getenv):
        imgs = []
        m = mtvimages.ProcessImages(imgs, None, None)
        result = m.get_http_thumb_path("/tmp/test.jpg")
        self.assertIn(":8080", result)

    @patch("mtvimages.ProcessImages.thumb_dir_check")
    @patch("mtvimages.ProcessImages.create_thumbnail")
    @patch("mtvimages.os.stat")
    @patch("mtvimages.ProcessImages.get_img_id", return_value="abc")
    @patch("mtvimages.ProcessImages.get_name", return_value="test")
    @patch("mtvimages.ProcessImages.get_thumb_path", return_value="/tmp/thumb.jpg")
    @patch("mtvimages.ProcessImages.get_http_thumb_path", return_value="localhost:8080/thumbnails/test.jpg")
    @patch("mtvimages.logging.error")
    def test_process(self, mock_log, mock_http, mock_thumb, mock_name, mock_imgid, mock_stat, mock_create, mock_dir):
        imgs = ["/tmp/test.jpg"]
        mock_stat.return_value.st_size = 123
        m = mtvimages.ProcessImages(imgs, MagicMock(), MagicMock())
        m.cursor.execute = MagicMock()
        m.conn.commit = MagicMock()
        m.process()

class TestProcessTVShowImages(unittest.TestCase):
    def test_class_init(self):
        try:
            imgs = []
            m = mtvimages.ProcessTVShowImages(imgs)
        except Exception as e:
            self.fail(f"ProcessTVShowImages() raised: {e}")

    @patch("mtvimages.os.path.exists", return_value=True)
    def test_tv_thumb_dir_check(self, mock_exists):
        imgs = []
        m = mtvimages.ProcessTVShowImages(imgs)
        m.tv_thumb_dir_check()

    @patch("mtvimages.Image.open")
    @patch("mtvimages.os.path.splitext", side_effect=lambda x: (x, ".jpg"))
    @patch("mtvimages.os.path.split", side_effect=lambda x: ("/tmp", "test"))
    @patch("mtvimages.os.path.join", side_effect=lambda a, b: f"{a}/{b}.jpg")
    @patch("mtvimages.os.getenv", return_value="/tmp")
    def test_create_thumbnail(self, mock_getenv, mock_join, mock_split, mock_ext, mock_open):
        imgs = []
        m = mtvimages.ProcessTVShowImages(imgs)
        mock_thumb = MagicMock()
        mock_open.return_value = mock_thumb
        mock_thumb.thumbnail.return_value = None
        mock_thumb.save.return_value = None
        result = m.create_thumbnail("/tmp/test.jpg")
        self.assertIn("/tmp", result)

    @patch("mtvimages.ProcessTVShowImages.tv_thumb_dir_check")
    @patch("mtvimages.ProcessTVShowImages.create_thumbnail")
    @patch("mtvimages.logging.info")
    def test_process_tv_thumbs(self, mock_log, mock_create, mock_dir):
        imgs = ["/tmp/test.jpg"]
        m = mtvimages.ProcessTVShowImages(imgs)
        m.process_tv_thumbs()

if __name__ == "__main__":
    unittest.main()
