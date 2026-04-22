
import unittest
from unittest.mock import patch, MagicMock
import mtvserverutils

class TestMedia(unittest.TestCase):
    def setUp(self):
        patcher_conn = patch("mtvserverutils.sqlite3.connect")
        patcher_cursor = patch("mtvserverutils.sqlite3.Cursor")
        self.addCleanup(patcher_conn.stop)
        self.mock_connect = patcher_conn.start()
        self.mock_conn = MagicMock()
        self.mock_cursor = MagicMock()
        self.mock_connect.return_value = self.mock_conn
        self.mock_conn.cursor.return_value = self.mock_cursor
        self.mock_cursor.description = [("col1",), ("col2",)]
        self.mock_cursor.fetchall.return_value = [(1, 2)]
        self.media = mtvserverutils.Media()

    def test_class_init(self):
        self.assertIsInstance(self.media, mtvserverutils.Media)

    def test_fetch_all_as_dict(self):
        self.mock_cursor.description = [("col1",), ("col2",)]
        self.mock_cursor.fetchall.return_value = [(1, 2)]
        result = self.media._fetch_all_as_dict()
        self.assertEqual(result, [{"col1": 1, "col2": 2}])

    def test_search(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"name": "test"}])
        result = self.media.search("test")
        self.assertIsInstance(result, list)

    def test_action(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "Action"}])
        result = self.media.action()
        self.assertIsInstance(result, list)

    def test_arnold(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "Arnold"}])
        result = self.media.arnold()
        self.assertIsInstance(result, list)

    def test_avatar(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "Avatar"}])
        result = self.media.avatar()
        self.assertIsInstance(result, list)

    def test_brucelee(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "BruceLee"}])
        result = self.media.brucelee()
        self.assertIsInstance(result, list)

    def test_brucewillis(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "BruceWillis"}])
        result = self.media.brucewillis()
        self.assertIsInstance(result, list)

    def test_buzz(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "Buzz"}])
        result = self.media.buzz()
        self.assertIsInstance(result, list)

    def test_cartoons(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "Cartoons"}])
        result = self.media.cartoons()
        self.assertIsInstance(result, list)

    def test_charliebrown(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "CharlieBrown"}])
        result = self.media.charliebrown()
        self.assertIsInstance(result, list)

    def test_comedy(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "Comedy"}])
        result = self.media.comedy()
        self.assertIsInstance(result, list)

    def test_cheechandchong(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "CheechAndChong"}])
        result = self.media.cheechandchong()
        self.assertIsInstance(result, list)

    def test_chucknorris(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "ChuckNorris"}])
        result = self.media.chucknorris()
        self.assertIsInstance(result, list)

    def test_clinteastwood(self):
        self.media._fetch_all_as_dict = MagicMock(return_value=[{"catagory": "ClintEastwood"}])
        result = self.media.clinteastwood()
        self.assertIsInstance(result, list)

    # ...repeat for all other methods, using the same pattern...

if __name__ == "__main__":
    unittest.main()
