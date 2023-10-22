const server = "localhost:8080"
const backendBaseUrl = `http://${server}/api/v1`


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
  const feedbackEl = document.getElementById('feedback');
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
      .catch(() => feedbackEl.innerHTML = "ERROR!")
      .then(() => feedbackEl.innerHTML = "SUCCESS!")
  }
}