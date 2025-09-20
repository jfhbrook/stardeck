const SECOND = 1000;

const DBUS_DESTINATION = 'org.jfhbrook.stardeck';
const DBUS_PATH = '/';
const DBUS_INTERFACE = 'org.jfhbrook.stardeck';

const SET_WINDOW_EVERY = 0.2 * SECOND;
const RESET_WINDOW_EVERY = 10 * SECOND;

function callback(methodName, ...args) {
  callDBus(DBUS_DESTINATION, DBUS_PATH, DBUS_INTERFACE, methodName, ...args);
}

function setInterval(callback, interval) {
  const timer = new QTimer();
  timer.interval = interval;
  timer.timeout.connect(callback);
  timer.start();
}

let caption = '';

function setWindow() {
  if (caption === workspace.activeWindow.caption) {
    return;
  }
  caption = workspace.activeWindow.caption;
  callback('SetWindow', caption);
}

function resetWindow() {
  caption = '';
}

setInterval(setWindow, SET_WINDOW_EVERY);
setInterval(resetWindow, RESET_WINDOW_EVERY);
