cd /home/pimedia/MTVP/go
rm ./mtvp
git pull
go build

if [ ! -f /etc/systemd/system/mtvp.service ]; then
    cp /home/pimedia/MTVP/go/mtvp.service /etc/systemd/system/mtvp.service
    systemctl daemon-reload
    systemctl enable mtvp.service
    systemctl start mtvp.service
fi

if [ -f /etc/systemd/system/mtvp.service ]; then
    systemctl start mtvp.service
fi