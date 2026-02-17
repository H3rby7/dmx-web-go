const server = "localhost:8080";
const backendBaseUrl = `http://${server}/mock`;

/**
 * Render feedback to 'feedback' element
 *
 * @param {string} content
 */
function feedback(content) {
  document.getElementById("feedback").innerHTML = content;
}

/**
 * Send a DMX changeset to the backend
 *
 * @param {[{ channel: number, value: byte }]} list of channels with values
 *
 * Example useage `sendDMX([{channelA: valueA, channelB: valueB}])`
 */
function sendDMX(list = []) {
  const url = `${backendBaseUrl}/api/read`;
  fetch(url, {
    headers: {
      "content-type": "application/json",
    },
    body: JSON.stringify({ List: list }),
    method: "PATCH",
  })
    .catch(() => feedback("ERROR! ERROR! ERROR! ERROR! ERROR! ERROR! ERROR! ERROR! ERROR! ERROR!"))
    .then(() => feedback(""));
}

/**
 * Register onChange function that sends the 'value' from the given HTML element
 * to the backend for the specified channel
 *
 * @param {string} elementId ID of the HTML element holding the channel
 * @param {int} channel for the dmx value
 *
 * Example useage `registerOnChangeSendDMX('elementA', 1)"`
 */
function registerOnChangeSendDMX(elementId, channel) {
  console.debug(`Registering onChange handler for #${elementId}`);
  const el = document.getElementById(elementId);
  if (!el) {
    console.error(`Element #${elementId} not found...`);
  }
  el.addEventListener("input", () => {
    const value = parseInt(el.value);
    sendDMX([{ channel, value }]);
  });
}
