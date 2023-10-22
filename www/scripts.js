const server = "localhost:8080"
const backendBaseUrl = `http://${server}/api/v1`

/**
 * Render feedback to 'feedback' element
 * 
 * @param {string} content 
 */
function feedback(content) {
  document.getElementById('feedback').innerHTML = content;
}

/**
 * Returns a function that, upon invocation, reads the value of 'trigger' 
 * and writes it to the innerHtml of 'display'.
 * 
 * @param {string} trigger ID of the HTML element holding the channel
 * @param {string} display  ID of the HTML element holding the dmx value
 * @returns the function to be used as event handler
 * 
 * Example useage `oninput="sendSingleDMX('triggerEl', 'displayEl')()"`
 */
function onChangeUpdate(trigger, display) {
  return () => {
    const val = document.getElementById(trigger).value;
    document.getElementById(display).innerHTML = val;
  }
}

/**
 * Returns a function that, upon invocation, reads 'channel' and 'value' from the given HTML elements 
 * and sends them to the backend.
 * 
 * @param {string} channelElementId ID of the HTML element holding the channel
 * @param {string} valueElementId  ID of the HTML element holding the dmx value
 * @returns the function to be used as event handler
 * 
 * Example useage `onclick="sendSingleDMX('elementA', 'elementB')()"`
 */
function sendSingleDMX(channelElementId, valueElementId) {
  const channelEl = document.getElementById(channelElementId)
  const valueEl = document.getElementById(valueElementId)
  const url = `${backendBaseUrl}/dmx`;
  return () => {
    const channel = parseInt(channelEl.value);
    const value = parseInt(valueEl.value);
    fetch(url, {
      headers: {
        "content-type": "application/json"
      },
      body: JSON.stringify({ List: [{ channel, value }] }),
      method: "PATCH"
    })
      .catch(() => feedback("ERROR!"))
      .then(() => feedback("SUCCESS!"))
  }
}

function fadeMultipleDMX(dmxList, fadeTimeMillis = 1) {
  const url = `${backendBaseUrl}/dmx/fade`;
  fetch(url, {
    headers: {
      "content-type": "application/json"
    },
    body: JSON.stringify({ fadeTimeMillis: fadeTimeMillis, scene: {list: dmxList }}),
    method: "PATCH"
  })
    .catch(() => feedback("ERROR!"))
    .then(() => feedback("SUCCESS!"))
}

// Channels used for dimming lights
const dimmerChannels = [1, 2, 3, 4, 8];

const sceneA = {
  fadeable: [
    { channel: 1, value: 50 },
    { channel: 2, value: 100 },
    { channel: 3, value: 150 },
  ]
};
const sceneB = {
  fadeable: [
    { channel: 1, value: 150 },
    { channel: 4, value: 200 },
  ]
};
const sceneC = {
  fadeable: [
    { channel: 8, value: 50 },
  ],
  settable: [
    { channel: 5, value: 255 },
    { channel: 6, value: 255 },
    { channel: 7, value: 255 },
  ]
};

function switchToScene(scene) {
  if (scene.settable) {
    // Send any DMX values that we do not want to fade-in as first request
    fadeMultipleDMX(scene.settable);
  }
  const chMap = new Map();
  for (let i = 0; i < dimmerChannels.length; i++) {
    chMap.set(dimmerChannels[i], 0);
  }
  for (let i = 0; i < scene.fadeable.length; i++) {
    const e = scene.fadeable[i];
    chMap.set(e.channel, e.value);
  }
  const dimmerValues = Array.from(chMap, ([channel, value]) => ({ channel, value }));
  fadeMultipleDMX(dimmerValues, 2000);
}