// Attach click listeners to all .mov-img images and log their id
document.addEventListener('DOMContentLoaded', function() {
    const images = document.querySelectorAll('.mov-img');
    images.forEach(function(img) {
        img.addEventListener('click', function() {
            const movId = img.id;
            console.log('Clicked image id:', movId);

            // Connect to WebSocket server (assume ws://localhost:8765/)
            const ws = new WebSocket('ws://10.0.4.41:8765/');

            ws.onopen = function() {
                // Send set_media command with mov-id
                ws.send(JSON.stringify({ command: 'set_media', media_id: movId }));
                // After set_media, send play command
                ws.send(JSON.stringify({ command: 'play' }));
            };

            ws.onmessage = function(event) {
                console.log('WebSocket message:', event.data);
            };

            ws.onerror = function(error) {
                console.error('WebSocket error:', error);
            };

            ws.onclose = function() {
                console.log('WebSocket connection closed');
            };
        });
    });
});

// document.addEventListener('DOMContentLoaded', () => {
const btnBack = document.getElementById('btn-back');
const btnPlay = document.getElementById('btn-play');
const btnPause = document.getElementById('btn-pause');
const btnNext = document.getElementById('btn-next');
// Add click event listeners
if (btnBack) {
    btnBack.addEventListener('click', () => {
        console.log('Back button clicked');
        // Add your custom logic here (e.g., previous page, previous track)
    });
}
if (btnPlay) {
    btnPlay.addEventListener('click', () => {
        console.log('Play button clicked');
        // Add your play logic here
    });
}
if (btnPause) {
    btnPause.addEventListener('click', () => {
        console.log('Pause button clicked');
        // Add your pause logic here
    });
}
if (btnNext) {
    btnNext.addEventListener('click', () => {
        console.log('Next button clicked');
        // Add your custom logic here (e.g., next page, next track)
    });
}

