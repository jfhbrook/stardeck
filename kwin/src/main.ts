const SECOND = 1000;

const DBUS_DESTINATION = 'org.jfhbrook.stardeck';
const DBUS_PATH = '/';
const DBUS_INTERFACE = 'org.jfhbrook.stardeck';

const WINDOW_INTERVAL = 1 * SECOND;

function callback(methodName: string, ...args: any[]) {
  callDBus(DBUS_DESTINATION, DBUS_PATH, DBUS_INTERFACE, methodName, ...args);
}

function setWindow() {
  callback('SetWindow', workspace.activeWindow.caption);
}

setInterval(setWindow, WINDOW_INTERVAL);
