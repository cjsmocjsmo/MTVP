var wsAddr = 'ws://10.0.4.41:8090/ws';

// Attach click listeners to all .mov-img images for set_media WebSocket command
document.addEventListener('DOMContentLoaded', function() {
    const movImgs = document.querySelectorAll('.mov-img');
    movImgs.forEach(function(img) {
        img.addEventListener('click', function() {
            const movId = img.id;
            console.log('[mtv.js] mov-img clicked. Selected movId:', movId);
            const ws = new WebSocket(wsAddr);
            ws.onopen = function() {
                console.log('[mtv.js] WebSocket opened for movId:', movId);
                ws.send(JSON.stringify({ command: 'set_media', media_id: movId }));
            };
            ws.onmessage = function(event) {
                console.log('[mtv.js] WebSocket message:', event.data);
            };
            ws.onerror = function(error) {
                console.error('[mtv.js] WebSocket error:', error);
            };
            ws.onclose = function() {
                console.log('[mtv.js] WebSocket connection closed for movId:', movId);
            };
        });
    });

    // Attach click listeners to all .mov-video-name elements for home_set_media WebSocket command
    const movVideoNames = document.querySelectorAll('.mov-video-name');
    movVideoNames.forEach(function(p) {
        p.addEventListener('click', function() {
            const vidId = p.id;
            console.log('[mtv.js] mov-video-name clicked. Selected VidId:', vidId);
            const ws = new WebSocket(wsAddr);
            ws.onopen = function() {
                console.log('[mtv.js] WebSocket opened for VidId:', vidId);
                ws.send(JSON.stringify({ command: 'home_set_media', media_id: vidId }));
            };
            ws.onmessage = function(event) {
                console.log('[mtv.js] WebSocket message:', event.data);
            };
            ws.onerror = function(error) {
                console.error('[mtv.js] WebSocket error:', error);
            };
            ws.onclose = function() {
                console.log('[mtv.js] WebSocket connection closed for VidId:', vidId);
            };
        });
    });

    // Attach click listeners to all TV episode buttons for tv_set_media WebSocket command
    const tvEpisodeBtns = document.querySelectorAll('.epi-div-item');
    tvEpisodeBtns.forEach(function(btn) {
        btn.addEventListener('click', function() {
            const tvId = btn.id;
            console.log('[mtv.js] TV episode button clicked. Selected TvId:', tvId);
            const ws = new WebSocket(wsAddr);
            ws.onopen = function() {
                console.log('[mtv.js] WebSocket opened for TvId:', tvId);
                ws.send(JSON.stringify({ command: 'tv_set_media', media_id: tvId }));
            };
            ws.onmessage = function(event) {
                console.log('[mtv.js] WebSocket message:', event.data);
            };
            ws.onerror = function(error) {
                console.error('[mtv.js] WebSocket error:', error);
            };
            ws.onclose = function() {
                console.log('[mtv.js] WebSocket connection closed for TvId:', tvId);
            };
        });
    });
// });
    // --- Movie Search Handler ---
// document.addEventListener('DOMContentLoaded', function() {
    const movForm = document.getElementById('mov-search-form');
    const movInput = document.getElementById('mov-search-input');
    const movResults = document.querySelector('.mov-search-results');
    if (movForm && movInput && movResults) {
        movForm.addEventListener('submit', function(e) {
            e.preventDefault();
            const query = movInput.value.trim();
            if (!query) return;
            fetch(`/movsearch?q=${encodeURIComponent(query)}`)
                .then(res => res.json())
                .then(data => renderMovieResults(data.results || [], movResults))
                .catch(() => { movResults.innerHTML = '<div class="no-results">Search error.</div>'; });
        });
    }

    // --- TV Search Handler ---
    const tvForm = document.getElementById('tv-search-form');
    const tvInput = document.getElementById('tv-search-input');
    const tvResults = document.querySelector('.tv-search-results');
    if (tvForm && tvInput && tvResults) {
        tvForm.addEventListener('submit', function(e) {
            e.preventDefault();
            const query = tvInput.value.trim();
            if (!query) return;
            fetch(`/tvsearch?q=${encodeURIComponent(query)}`)
                .then(res => res.json())
                .then(data => renderTVResults(data.results || [], tvResults))
                .catch(() => { tvResults.innerHTML = '<div class="no-results">Search error.</div>'; });
        });
    }
});


function renderMovieResults(results, container) {
    console.log('[mtv.js] Enter renderMovieResults, results:', results, 'container:', container);
    container.innerHTML = '';
    if (!results.length) {
        console.log('[mtv.js] No movies found in renderMovieResults');
        container.innerHTML = '<div class="no-results">No movies found.</div>';
        return;
    }
    const div = document.createElement('div');
    div.className = 'result-item';
    results.forEach(r => {
        if (r.HttpThumbPath) {
            const img = document.createElement('img');
            img.id = r.MovId || '';
            img.src = r.HttpThumbPath;
            img.alt = r.Name || '';
            // Add click event to send set-media command to ws server
            img.addEventListener('click', function() {
                const movId = img.id;
                console.log('[mtv.js] Movie image clicked. Selected movId:', movId);
                // Connect to WebSocket server (adjust ws:// address as needed)
                const ws = new WebSocket(wsAddr);
                ws.onopen = function() {
                    console.log('[mtv.js] WebSocket opened for movId:', movId);
                    ws.send(JSON.stringify({ command: 'set_media', media_id: movId }));
                };
                ws.onmessage = function(event) {
                    console.log('[mtv.js] WebSocket message:', event.data);
                };
                ws.onerror = function(error) {
                    console.error('[mtv.js] WebSocket error:', error);
                };
                ws.onclose = function() {
                    console.log('[mtv.js] WebSocket connection closed for movId:', movId);
                };
            });
            div.appendChild(img);
        }
    });
    container.appendChild(div);
    console.log('[mtv.js] Exit renderMovieResults');
}

function escapeHtml(text) {
    console.log('[mtv.js] Enter escapeHtml, text:', text);
    const result = String(text).replace(/[&<>"']/g, function(m) {
        return ({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;','\'':'&#39;'}[m]);
    });
    console.log('[mtv.js] Exit escapeHtml, result:', result);
    return result;
}

function renderUpdateResults(payload, container, mediaLabel, summaryKey, insertedKey) {
    if (!container) {
        return;
    }

    const summary = payload && payload.summary ? payload.summary : {};
    const batch = summary[summaryKey] || {};
    const errors = batch.errors || [];
    const inserted = insertedKey ? payload[insertedKey] : batch.inserted;

    let html = '<div class="no-results">';
    html += `<strong>${escapeHtml(mediaLabel)} update complete.</strong><br>`;
    html += `Status: ${escapeHtml(payload.status || 'unknown')}<br>`;
    html += `Scanned: ${Number(batch.scanned || 0)}<br>`;
    html += `Inserted: ${Number(inserted || 0)}<br>`;
    html += `Skipped: ${Number(batch.skipped || 0)}<br>`;
    html += `Failed: ${Number(batch.failed || 0)}`;

    if (errors.length) {
        html += '<br><br><strong>Errors</strong><br>';
        html += errors.map((err) => `<div>${escapeHtml(err)}</div>`).join('');
    }

    html += '</div>';
    container.innerHTML = html;
}

document.addEventListener('DOMContentLoaded', function() {
    const movUpdateButton = document.getElementById('mov-update-button');
    const movResults = document.querySelector('.mov-search-results');
    if (movUpdateButton && movResults) {
        movUpdateButton.addEventListener('click', function(event) {
            event.preventDefault();
            movResults.innerHTML = '<div class="no-results">Running movie update...</div>';

            fetch('/movupdate')
                .then((res) => res.json())
                .then((data) => renderUpdateResults(data, movResults, 'Movie', 'movies', 'moviesInserted'))
                .catch(() => {
                    movResults.innerHTML = '<div class="no-results">Movie update failed.</div>';
                });
        });
    }

    const tvUpdateButton = document.getElementById('tv-update-button');
    const tvResults = document.querySelector('.tv-search-results');

    if (tvUpdateButton && tvResults) {
        tvUpdateButton.addEventListener('click', function(event) {
            event.preventDefault();
            tvResults.innerHTML = '<div class="no-results">Running TV update...</div>';

            fetch('/tvupdate')
                .then((res) => res.json())
                .then((data) => renderUpdateResults(data, tvResults, 'TV', 'tvshows', 'tvshowsInserted'))
                .catch(() => {
                    tvResults.innerHTML = '<div class="no-results">TV update failed.</div>';
                });
        });
    }
});

// document.addEventListener('DOMContentLoaded', () => {
const btnBack = document.getElementById('btn-back');
const btnStop = document.getElementById('btn-stop');
const btnPlay = document.getElementById('btn-play');
const btnPause = document.getElementById('btn-pause');
const btnNext = document.getElementById('btn-next');
// Add click event listeners
if (btnBack) {
    btnBack.addEventListener('click', () => {
        console.log('Back button clicked');
        // Send previous command to WebSocket server
        const ws = new WebSocket(wsAddr);
        ws.onopen = function() {
            ws.send(JSON.stringify({ command: 'previous' }));
            console.log('[mtv.js] Sent previous command to server');
            ws.close();
        };
        ws.onerror = function(error) {
            console.error('[mtv.js] WebSocket error (previous):', error);
        };
    });
}
if (btnStop) {
    btnStop.addEventListener('click', () => {
        console.log('Stop button clicked');
        // Send stop command to WebSocket server
        const ws = new WebSocket(wsAddr);
        ws.onopen = function() {
            ws.send(JSON.stringify({ command: 'stop' }));
            console.log('[mtv.js] Sent stop command to server');
            ws.close();
        };
        ws.onerror = function(error) {
            console.error('[mtv.js] WebSocket error (stop):', error);
        };
    });
}
if (btnPlay) {
    btnPlay.addEventListener('click', () => {
        console.log('Play button clicked');
        // Send play command to WebSocket server
        const ws = new WebSocket(wsAddr);
        ws.onopen = function() {
            ws.send(JSON.stringify({ command: 'play' }));
            console.log('[mtv.js] Sent play command to server');
            ws.close();
        };
        ws.onerror = function(error) {
            console.error('[mtv.js] WebSocket error (play):', error);
        };
    });
}
if (btnPause) {
    btnPause.addEventListener('click', () => {
        console.log('Pause button clicked');
        // Send pause command to WebSocket server
        const ws = new WebSocket(wsAddr);
        ws.onopen = function() {
            ws.send(JSON.stringify({ command: 'pause' }));
            console.log('[mtv.js] Sent pause command to server');
            ws.close();
        };
        ws.onerror = function(error) {
            console.error('[mtv.js] WebSocket error (pause):', error);
        };
    });
}
if (btnNext) {
    btnNext.addEventListener('click', () => {
        console.log('Next button clicked');
        // Send next command to WebSocket server
        const ws = new WebSocket(wsAddr);
        ws.onopen = function() {
            ws.send(JSON.stringify({ command: 'next' }));
            console.log('[mtv.js] Sent next command to server');
            ws.close();
        };
        ws.onerror = function(error) {
            console.error('[mtv.js] WebSocket error (next):', error);
        };
    });
}

document.addEventListener('DOMContentLoaded', function() {
    const radarImage = document.getElementById('radar-image');
    if (!radarImage) {
        return;
    }

    const radarPage = document.querySelector('.radar-page');
    const updatedLabel = document.getElementById('radar-last-updated');
    const refreshMs = Number(radarPage?.dataset?.radarRefreshMs) || 300000;
    const baseURL = radarImage.dataset.radarBase || radarImage.src;

    function setLastUpdatedText() {
        if (!updatedLabel) {
            return;
        }
        const now = new Date();
        updatedLabel.textContent = 'Last updated: ' + now.toLocaleString();
    }

    function refreshRadarImage() {
        if (document.visibilityState !== 'visible') {
            return;
        }
        const separator = baseURL.includes('?') ? '&' : '?';
        radarImage.src = baseURL + separator + 'ts=' + Date.now();
        setLastUpdatedText();
    }

    refreshRadarImage();
    setInterval(refreshRadarImage, refreshMs);

    document.addEventListener('visibilitychange', function() {
        if (document.visibilityState === 'visible') {
            refreshRadarImage();
        }
    });
});