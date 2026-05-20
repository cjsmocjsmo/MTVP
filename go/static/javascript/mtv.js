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
// });