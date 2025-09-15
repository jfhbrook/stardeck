const SECOND = 1000;

const DBUS_DESTINATION = 'org.jfhbrook.stardeck';
const DBUS_PATH = '/';
const DBUS_INTERFACE = 'org.jfhbrook.stardeck';

function callback(methodName, ...args) {
  callDBus(DBUS_DESTINATION, DBUS_PATH, DBUS_INTERFACE, methodName, ...args);
}

function setWindow() {
  callback('SetWindow', workspace.activeWindow.caption);
}

let timer = new QTimer();
timer.interval = 1 * SECOND;
timer.timeout.connect(setWindow);
timer.start();
