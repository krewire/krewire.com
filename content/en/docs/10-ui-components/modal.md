---
title: "Modal"
description: "Accessible modal dialog built on the native HTML dialog element."
date: "2026-11-09"
---

# Modal (`Modal.kiw`)

Category: **Overlays & Dialogs**  
Component File: `components/Modal.kiw`  
Source: `github.com/krewire/ui/components/Modal.kiw`

The `<Modal />` component renders an accessible modal dialog leveraging the browser's native `<dialog>` element. It features pop-in animations, backdrop blur, keyboard `Escape` dismissal, and focus trapping without third-party dependencies.

---

## 1. Copy & Own Installation

```bash
mkdir -p ./components
cp path/to/krewire/ui/components/Modal.kiw ./components/
```

---

## 2. Examples & Usage

```html
<!-- Trigger Button -->
<Button variant="primary" onClick="document.getElementById('confirm-modal').showModal()">
  Open Confirmation Modal
</Button>

<!-- Native Modal Dialog -->
<Modal id="confirm-modal" title="Confirm Workload Deployment">
  <div slot="body">
    <p>Are you sure you want to deploy <strong>v1.2.0</strong> to the production cluster?</p>
  </div>
  <div slot="actions">
    <Button variant="ghost" onClick="document.getElementById('confirm-modal').close()">
      Cancel
    </Button>
    <Button variant="primary" onClick="executeDeployment()">
      Deploy Now
    </Button>
  </div>
</Modal>
```

---

## 3. Accessibility & Keyboard Control

- **Escape Key**: Automatically closes the dialog using native browser events.
- **Backdrop Click**: Dismisses the modal when users click outside the dialog box.
- **Focus Isolation**: Focus is automatically confined inside the dialog while open.
