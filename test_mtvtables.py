
import unittest
from unittest.mock import patch, MagicMock
import mtvtables

class TestCreateTables(unittest.TestCase):
    def setUp(self):
        patcher_conn = patch("mtvtables.sqlite3.connect")
        self.addCleanup(patcher_conn.stop)
        self.mock_connect = patcher_conn.start()
        self.mock_conn = MagicMock()
        self.mock_cursor = MagicMock()
        self.mock_connect.return_value = self.mock_conn
        self.mock_conn.cursor.return_value = self.mock_cursor
        self.tables = mtvtables.CreateTables()

    def test_class_init(self):
        self.assertIsInstance(self.tables, mtvtables.CreateTables)

    def test_create_tables(self):
        self.tables.cursor.execute = MagicMock()
        self.tables.create_tables()
        self.assertTrue(self.tables.cursor.execute.called)

if __name__ == "__main__":
    unittest.main()
