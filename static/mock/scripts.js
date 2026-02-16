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
    .catch(() => feedback("ERROR!"))
    .then(() => feedback("SUCCESS!"));
}

const sceneA = [
  { channel: 1, value: 50 },
  { channel: 2, value: 100 },
  { channel: 3, value: 150 },
  { channel: 4, value: 0 },
  { channel: 5, value: 0 },
  { channel: 6, value: 0 },
  { channel: 7, value: 0 },
];
const sceneB = [
  { channel: 1, value: 150 },
  { channel: 2, value: 0 },
  { channel: 3, value: 0 },
  { channel: 4, value: 200 },
  { channel: 5, value: 0 },
  { channel: 6, value: 0 },
  { channel: 7, value: 0 },
];
const sceneC = [
  { channel: 1, value: 0 },
  { channel: 2, value: 0 },
  { channel: 3, value: 0 },
  { channel: 4, value: 0 },
  { channel: 5, value: 255 },
  { channel: 6, value: 255 },
  { channel: 7, value: 255 },
];
