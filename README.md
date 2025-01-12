# Windows MessageBox Flags

The `MessageBoxW` function in the Windows API allows you to display customizable message boxes with various icons, buttons, and behaviors. This guide will help you understand the different flags that can be used to personalize the appearance and functionality of message boxes in your Windows applications.

## Icon Flags

Choose an icon to display in the message box. Icons help convey the nature of the message and grab the user’s attention.

| Flag Value | Constant                 | Description                               | Icon   |
|------------|--------------------------|-------------------------------------------|--------|
| `0x10`     | `MB_ICONHAND`            | A **hand** or **stop** icon               | ❌     |
| `0x20`     | `MB_ICONQUESTION`        | A **question mark** icon                  | ❓     |
| `0x30`     | `MB_ICONEXCLAMATION`     | A **warning** icon                       | ⚠️     |
| `0x40`     | `MB_ICONINFORMATION`     | An **information** icon                  | ℹ️     |

These icons help users quickly identify the type of message being presented to them. For example, use a warning icon when there’s a problem, or an info icon for helpful details.

---

## Button Flags

Control the buttons that appear in the message box. You can offer different choices for the user depending on the context.

| Flag Value | Constant                   | Description                                      |
|------------|----------------------------|--------------------------------------------------|
| `0x0`      | `MB_OK`                    | Displays an **OK** button only.                  |
| `0x1`      | `MB_OKCANCEL`              | Displays **OK** and **Cancel** buttons.          |
| `0x2`      | `MB_ABORTRETRYIGNORE`      | Displays **Abort**, **Retry**, and **Ignore** buttons. |
| `0x3`      | `MB_YESNOCANCEL`           | Displays **Yes**, **No**, and **Cancel** buttons. |
| `0x4`      | `MB_YESNO`                 | Displays **Yes** and **No** buttons.             |
| `0x5`      | `MB_RETRYCANCEL`           | Displays **Retry** and **Cancel** buttons.       |

### Default Focus on Buttons
You can also specify which button should be the default when the user opens the message box. This helps guide the user’s choice.

| Flag Value | Constant       | Description                                          |
|------------|----------------|------------------------------------------------------|
| `0x0`      | `MB_DEFBUTTON1`| The first button is the default.                     |
| `0x100`    | `MB_DEFBUTTON2`| The second button is the default.                    |
| `0x200`    | `MB_DEFBUTTON3`| The third button is the default.                     |

---

## Modal Behavior Flags

You can control how the message box behaves with respect to the application or system. These flags determine how the message box interacts with the rest of the system.

| Flag Value | Constant         | Description                                           |
|------------|------------------|-------------------------------------------------------|
| `0x0`      | `MB_APPLMODAL`   | Blocks input to the calling application only.         |
| `0x1000`   | `MB_SYSTEMMODAL` | Blocks input to all applications on the system.       |
| `0x2000`   | `MB_TASKMODAL`   | Blocks input to other apps in the same task.          |

This allows you to choose how restrictive the message box should be, depending on the level of importance of the message.

---

## 📝 Example Usage

```go
syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
    0,
    uintptr(unsafe.Pointer(text)),
    uintptr(unsafe.Pointer(title)),
    0x40 | 0x1, // MB_ICONINFORMATION + MB_OKCANCEL
)
