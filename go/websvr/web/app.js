document.addEventListener('DOMContentLoaded', function() {
    initializeNavigation();
    initializeCharts();
    initializeDashboard();
    initializeDeviceManagement();
    initializeTopology();
    startRealTimeUpdates();
});

function initializeNavigation() {
    const navItems = document.querySelectorAll('.nav-item');
    const pages = document.querySelectorAll('.page');
    const menuToggle = document.getElementById('menuToggle');
    const sidebar = document.getElementById('sidebar');

    navItems.forEach(item => {
        item.addEventListener('click', function(e) {
            e.preventDefault();

            navItems.forEach(nav => nav.classList.remove('active'));
            this.classList.add('active');

            const targetPage = this.dataset.page;
            pages.forEach(page => {
                page.classList.remove('active');
                if (page.id === targetPage) {
                    page.classList.add('active');
                }
            });

            if (window.innerWidth <= 768) {
                sidebar.classList.remove('active');
            }
        });
    });

    if (menuToggle) {
        menuToggle.addEventListener('click', function() {
            sidebar.classList.toggle('active');
        });
    }

    document.addEventListener('click', function(e) {
        if (window.innerWidth <= 768) {
            if (!sidebar.contains(e.target) && !menuToggle.contains(e.target)) {
                sidebar.classList.remove('active');
            }
        }
    });
}

function initializeCharts() {
    const trafficCtx = document.getElementById('trafficChart');
    if (trafficCtx) {
        const trafficChart = new Chart(trafficCtx, {
            type: 'line',
            data: {
                labels: generateTimeLabels(24),
                datasets: [{
                    label: 'Inbound Traffic',
                    data: generateRandomData(24, 500, 2000),
                    borderColor: '#3b82f6',
                    backgroundColor: 'rgba(59, 130, 246, 0.1)',
                    tension: 0.3,
                    fill: true
                }, {
                    label: 'Outbound Traffic',
                    data: generateRandomData(24, 300, 1500),
                    borderColor: '#10b981',
                    backgroundColor: 'rgba(16, 185, 129, 0.1)',
                    tension: 0.3,
                    fill: true
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        position: 'top',
                    },
                    tooltip: {
                        mode: 'index',
                        intersect: false,
                    }
                },
                scales: {
                    y: {
                        beginAtZero: true,
                        ticks: {
                            callback: function(value) {
                                return value + ' Mbps';
                            }
                        }
                    }
                }
            }
        });
    }

    const bandwidthCtx = document.getElementById('bandwidthChart');
    if (bandwidthCtx) {
        const bandwidthChart = new Chart(bandwidthCtx, {
            type: 'doughnut',
            data: {
                labels: ['HTTP/HTTPS', 'FTP', 'Email', 'VoIP', 'Video Streaming', 'Other'],
                datasets: [{
                    data: [35, 10, 15, 12, 20, 8],
                    backgroundColor: [
                        '#3b82f6',
                        '#10b981',
                        '#f59e0b',
                        '#ef4444',
                        '#8b5cf6',
                        '#6b7280'
                    ]
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        position: 'right',
                    }
                }
            }
        });
    }

    const latencyCtx = document.getElementById('latencyHeatmap');
    if (latencyCtx) {
        const latencyChart = new Chart(latencyCtx, {
            type: 'bar',
            data: {
                labels: ['DC-A', 'DC-B', 'Branch-1', 'Branch-2', 'Branch-3', 'Remote-1'],
                datasets: [{
                    label: 'Average Latency',
                    data: [5, 8, 15, 22, 18, 45],
                    backgroundColor: function(context) {
                        const value = context.raw;
                        if (value < 10) return '#10b981';
                        if (value < 30) return '#f59e0b';
                        return '#ef4444';
                    }
                }]
            },
            options: {
                responsive: true,
                maintainAspectRatio: false,
                scales: {
                    y: {
                        beginAtZero: true,
                        ticks: {
                            callback: function(value) {
                                return value + ' ms';
                            }
                        }
                    }
                }
            }
        });
    }
}

function initializeDashboard() {
    updateNotificationCount();
    animateStatCards();
}

function updateNotificationCount() {
    const notificationCount = document.getElementById('notificationCount');
    if (notificationCount) {
        setInterval(() => {
            const count = Math.floor(Math.random() * 10);
            notificationCount.textContent = count;
            notificationCount.style.display = count > 0 ? 'block' : 'none';
        }, 30000);
    }
}

function animateStatCards() {
    const statNumbers = document.querySelectorAll('.stat-number');
    statNumbers.forEach(stat => {
        const finalValue = parseInt(stat.textContent) || parseFloat(stat.textContent);
        if (!isNaN(finalValue)) {
            animateValue(stat, 0, finalValue, 1500);
        }
    });
}

function animateValue(element, start, end, duration) {
    const isFloat = end % 1 !== 0;
    const range = end - start;
    const startTime = performance.now();

    function update(currentTime) {
        const elapsed = currentTime - startTime;
        const progress = Math.min(elapsed / duration, 1);

        const value = start + (range * easeOutQuart(progress));
        element.textContent = isFloat ? value.toFixed(2) + '%' : Math.floor(value);

        if (progress < 1) {
            requestAnimationFrame(update);
        }
    }

    requestAnimationFrame(update);
}

function easeOutQuart(t) {
    return 1 - Math.pow(1 - t, 4);
}

function initializeDeviceManagement() {
    const deviceSearch = document.getElementById('deviceSearch');
    if (deviceSearch) {
        deviceSearch.addEventListener('input', function() {
            const searchTerm = this.value.toLowerCase();
            const rows = document.querySelectorAll('#deviceTableBody tr');

            rows.forEach(row => {
                const text = row.textContent.toLowerCase();
                row.style.display = text.includes(searchTerm) ? '' : 'none';
            });
        });
    }

    // Clear the existing static rows first
    const deviceTableBody = document.getElementById('deviceTableBody');
    if (deviceTableBody) {
        deviceTableBody.innerHTML = '';
    }

    populateDeviceTable();
}

async function populateDeviceTable() {
    const deviceTableBody = document.getElementById('deviceTableBody');
    if (!deviceTableBody) return;

    let devices = [];

    try {
        // Try to fetch from REST API first
        const queryParams = new URLSearchParams({
            body: JSON.stringify({
                text: "select * from Device",
                rootType: "device",
                properties: ["*"],
                matchCase: true
            })
        });

        const response = await fetch(`/netop/0/dev-inv?${queryParams}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            }
        });

        if (response.ok) {
            const data = await response.json();
            if (data && data.list && Array.isArray(data.list)) {
                // Transform API response to match UI format
                devices = data.list.map(device => ({
                    id: device.id,
                    name: device.name,
                    type: mapDeviceType(device.type),
                    ip: device.ipAddress,
                    location: device.location,
                    status: mapDeviceStatus(device.status),
                    uptime: device.uptime,
                    model: device.model,
                    serial: device.serialNumber,
                    firmware: device.firmwareVersion,
                    mac: device.macAddress,
                    subnet: device.subnetMask,
                    gateway: device.gateway,
                    vlan: device.vlan
                }));
            } else {
                throw new Error('Invalid API response format');
            }
        } else {
            throw new Error(`API request failed with status ${response.status}`);
        }
    } catch (error) {
        console.warn('Failed to fetch devices from API, falling back to mock data:', error);
        // Fallback to mock data
        devices = getMockDevices();
    }

    // Clear existing rows
    deviceTableBody.innerHTML = '';

    // Populate table with devices
    devices.forEach(device => {
        const row = document.createElement('tr');
        row.setAttribute('data-device', JSON.stringify(device));
        row.style.cursor = 'pointer';

        row.innerHTML = `
            <td>${device.id}</td>
            <td>${device.name}</td>
            <td><span class="device-type">${device.type}</span></td>
            <td>${device.ip}</td>
            <td>${device.location}</td>
            <td><span class="status-badge ${device.status === 'online' ? 'online' : device.status === 'maintenance' ? 'warning' : 'offline'}">${device.status.charAt(0).toUpperCase() + device.status.slice(1)}</span></td>
            <td>${device.uptime}</td>
            <td class="action-buttons">
                <button class="btn-icon" title="Configure"><i class="fas fa-cog"></i></button>
                <button class="btn-icon" title="Monitor"><i class="fas fa-chart-line"></i></button>
                <button class="btn-icon" title="Reboot"><i class="fas fa-redo"></i></button>
            </td>
        `;

        // Add click handler to the row
        row.addEventListener('click', function(e) {
            // Don't open modal if clicking on action buttons
            if (!e.target.closest('.action-buttons')) {
                openDeviceModal(device);
            }
        });

        // Add click handlers to action buttons
        const buttons = row.querySelectorAll('.btn-icon');
        buttons[0].addEventListener('click', function(e) {
            e.stopPropagation();
            configureDevice(device.id);
        });
        buttons[1].addEventListener('click', function(e) {
            e.stopPropagation();
            monitorDevice(device.id);
        });
        buttons[2].addEventListener('click', function(e) {
            e.stopPropagation();
            rebootDevice(device.id);
        });

        deviceTableBody.appendChild(row);
    });
}

function mapDeviceType(apiType) {
    const typeMap = {
        'DEVICE_TYPE_SWITCH': 'Switch',
        'DEVICE_TYPE_ROUTER': 'Router',
        'DEVICE_TYPE_FIREWALL': 'Firewall',
        'DEVICE_TYPE_ACCESS_POINT': 'Access Point'
    };
    return typeMap[apiType] || apiType;
}

function mapDeviceStatus(apiStatus) {
    const statusMap = {
        'DEVICE_STATUS_ONLINE': 'online',
        'DEVICE_STATUS_OFFLINE': 'offline',
        'DEVICE_STATUS_WARNING': 'warning',
        'DEVICE_STATUS_MAINTENANCE': 'maintenance'
    };
    return statusMap[apiStatus] || 'offline';
}

function getMockDevices() {
    return [
        {
            id: 'SW-001',
            name: 'Core Switch A',
            type: 'Switch',
            ip: '192.168.1.10',
            location: 'Datacenter A',
            status: 'online',
            uptime: '45d 12h',
            model: 'Cisco Catalyst 9300',
            serial: 'FCW2148G0AB',
            firmware: 'IOS-XE 17.6.4',
            mac: '00:1B:54:C2:3A:4E',
            subnet: '255.255.255.0',
            gateway: '192.168.1.1',
            vlan: '10'
        },
        {
            id: 'SW-002',
            name: 'Core Switch B',
            type: 'Switch',
            ip: '192.168.1.11',
            location: 'Datacenter B',
            status: 'online',
            uptime: '30d 5h',
            model: 'Cisco Catalyst 9300',
            serial: 'FCW2149G0AB',
            firmware: 'IOS-XE 17.6.4',
            mac: '00:1B:54:C2:3A:4F',
            subnet: '255.255.255.0',
            gateway: '192.168.1.1',
            vlan: '10'
        },
        {
            id: 'RT-042',
            name: 'Edge Router 42',
            type: 'Router',
            ip: '192.168.2.1',
            location: 'Building B',
            status: 'warning',
            uptime: '12d 8h',
            model: 'Cisco ISR 4331',
            serial: 'FJC2148W0LA',
            firmware: 'IOS-XE 16.12.4',
            mac: '00:1A:2B:3C:4D:42',
            subnet: '255.255.255.0',
            gateway: '192.168.2.254',
            vlan: '20'
        },
        {
            id: 'RT-001',
            name: 'Main Router',
            type: 'Router',
            ip: '192.168.0.254',
            location: 'Datacenter A',
            status: 'online',
            uptime: '90d 12h',
            model: 'Cisco ISR 4451',
            serial: 'FJC2149W0LB',
            firmware: 'IOS-XE 16.12.5',
            mac: '00:1A:2B:3C:4D:5E',
            subnet: '255.255.255.0',
            gateway: '192.168.0.1',
            vlan: '1'
        },
        {
            id: 'FW-003',
            name: 'Main Firewall',
            type: 'Firewall',
            ip: '192.168.0.1',
            location: 'Datacenter A',
            status: 'online',
            uptime: '120d 5h',
            model: 'Fortinet FortiGate 200F',
            serial: 'FG200FTK20012346',
            firmware: 'FortiOS 7.2.5',
            mac: '00:09:0F:09:00:03',
            subnet: '255.255.255.0',
            gateway: '192.168.0.254',
            vlan: '1'
        },
        {
            id: 'AP-101',
            name: 'WiFi AP Floor 1',
            type: 'Access Point',
            ip: '192.168.10.101',
            location: 'Building A',
            status: 'online',
            uptime: '7d 3h',
            model: 'Cisco Aironet 2802i',
            serial: 'FCZ2049W0XY',
            firmware: 'AireOS 8.10.151.0',
            mac: '00:3A:9A:47:C5:B2',
            subnet: '255.255.255.0',
            gateway: '192.168.10.1',
            vlan: '100'
        },
        {
            id: 'FW-002',
            name: 'Backup Firewall',
            type: 'Firewall',
            ip: '192.168.0.2',
            location: 'Datacenter B',
            status: 'maintenance',
            uptime: 'N/A',
            model: 'Fortinet FortiGate 100F',
            serial: 'FG100FTK20012345',
            firmware: 'FortiOS 7.2.4',
            mac: '00:09:0F:09:00:02',
            subnet: '255.255.255.0',
            gateway: '192.168.0.1',
            vlan: '1'
        },
        {
            id: 'SW-015',
            name: 'Edge Switch 15',
            type: 'Switch',
            ip: '192.168.3.15',
            location: 'Building C',
            status: 'offline',
            uptime: 'N/A',
            model: 'Cisco Catalyst 2960X',
            serial: 'FOC2049Y0ZB',
            firmware: 'IOS 15.2(7)E3',
            mac: '00:1C:F6:C0:33:15',
            subnet: '255.255.255.0',
            gateway: '192.168.3.1',
            vlan: '30'
        }
    ];
}

function initializeTopology() {
    const canvas = document.getElementById('topologyCanvas');
    if (!canvas) return;

    const ctx = canvas.getContext('2d');
    canvas.width = canvas.offsetWidth;
    canvas.height = canvas.offsetHeight;

    const nodes = [
        { x: canvas.width / 2, y: canvas.height / 2, label: 'Core Switch', type: 'switch' },
        { x: canvas.width / 2 - 150, y: canvas.height / 2 - 100, label: 'Router 1', type: 'router' },
        { x: canvas.width / 2 + 150, y: canvas.height / 2 - 100, label: 'Router 2', type: 'router' },
        { x: canvas.width / 2 - 200, y: canvas.height / 2 + 100, label: 'Switch 1', type: 'switch' },
        { x: canvas.width / 2, y: canvas.height / 2 + 150, label: 'Switch 2', type: 'switch' },
        { x: canvas.width / 2 + 200, y: canvas.height / 2 + 100, label: 'Switch 3', type: 'switch' },
    ];

    const connections = [
        [0, 1], [0, 2], [0, 3], [0, 4], [0, 5],
        [1, 3], [2, 5], [3, 4], [4, 5]
    ];

    function drawTopology() {
        ctx.clearRect(0, 0, canvas.width, canvas.height);

        ctx.strokeStyle = '#e5e7eb';
        ctx.lineWidth = 2;
        connections.forEach(conn => {
            ctx.beginPath();
            ctx.moveTo(nodes[conn[0]].x, nodes[conn[0]].y);
            ctx.lineTo(nodes[conn[1]].x, nodes[conn[1]].y);
            ctx.stroke();
        });

        nodes.forEach(node => {
            ctx.fillStyle = node.type === 'router' ? '#3b82f6' : '#10b981';
            ctx.beginPath();
            ctx.arc(node.x, node.y, 25, 0, 2 * Math.PI);
            ctx.fill();

            ctx.fillStyle = 'white';
            ctx.font = '20px Font Awesome 5 Free';
            ctx.textAlign = 'center';
            ctx.textBaseline = 'middle';
            ctx.fillText(node.type === 'router' ? '\uf6ff' : '\uf233', node.x, node.y);

            ctx.fillStyle = '#111827';
            ctx.font = '12px sans-serif';
            ctx.fillText(node.label, node.x, node.y + 40);
        });
    }

    drawTopology();

    window.addEventListener('resize', () => {
        canvas.width = canvas.offsetWidth;
        canvas.height = canvas.offsetHeight;
        drawTopology();
    });
}

function startRealTimeUpdates() {
    setInterval(() => {
        // Only update dashboard elements if we're on the dashboard page
        const dashboardPage = document.getElementById('dashboard');
        if (dashboardPage && dashboardPage.classList.contains('active')) {
            updateRandomStats();
            updateDeviceStatuses();
            updateIncidentList();
        }
    }, 5000);
}

function updateRandomStats() {
    const activeDevices = document.querySelector('.stat-card:nth-child(1) .stat-number');
    if (activeDevices) {
        const current = parseInt(activeDevices.textContent);
        const change = Math.floor(Math.random() * 5) - 2;
        activeDevices.textContent = Math.max(240, Math.min(255, current + change));
    }

    const avgLatency = document.querySelector('.stat-card:nth-child(4) .stat-number');
    if (avgLatency) {
        const current = parseFloat(avgLatency.textContent);
        const change = (Math.random() * 2 - 1).toFixed(0);
        avgLatency.textContent = Math.max(8, Math.min(20, current + parseFloat(change))) + 'ms';
    }
}

function updateDeviceStatuses() {
    const statusBars = document.querySelectorAll('.status-fill');
    statusBars.forEach(bar => {
        const currentWidth = parseFloat(bar.style.width);
        const change = (Math.random() * 4 - 2);
        const newWidth = Math.max(80, Math.min(100, currentWidth + change));
        bar.style.width = newWidth + '%';

        const statusValue = bar.parentElement.nextElementSibling;
        if (statusValue) {
            statusValue.textContent = Math.floor(newWidth) + '% Online';
        }
    });
}

function updateIncidentList() {
    const incidentTimes = document.querySelectorAll('.incident-time');
    incidentTimes.forEach(time => {
        const text = time.textContent;
        const match = text.match(/(\d+)/);
        if (match) {
            const value = parseInt(match[1]);
            if (text.includes('minute')) {
                time.textContent = (value + 1) + ' minutes ago';
            }
        }
    });
}

function refreshDashboard() {
    const refreshBtn = document.querySelector('.refresh-btn button i');
    if (refreshBtn) {
        refreshBtn.classList.add('fa-spin');
        setTimeout(() => {
            refreshBtn.classList.remove('fa-spin');
            showNotification('Dashboard refreshed successfully');
        }, 1000);
    }
}

function addDevice() {
    showNotification('Device configuration dialog would open here');
}

function zoomIn() {
    showNotification('Zooming in on topology');
}

function zoomOut() {
    showNotification('Zooming out on topology');
}

function resetView() {
    showNotification('Resetting topology view');
}

function showNotification(message) {
    const notification = document.createElement('div');
    notification.className = 'notification';
    notification.textContent = message;
    notification.style.cssText = `
        position: fixed;
        bottom: 20px;
        right: 20px;
        background: #111827;
        color: white;
        padding: 12px 20px;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0,0,0,0.1);
        z-index: 9999;
        animation: slideIn 0.3s ease;
    `;

    document.body.appendChild(notification);

    setTimeout(() => {
        notification.style.animation = 'slideOut 0.3s ease';
        setTimeout(() => {
            document.body.removeChild(notification);
        }, 300);
    }, 3000);
}

function generateTimeLabels(count) {
    const labels = [];
    for (let i = count - 1; i >= 0; i--) {
        labels.push(i + ':00');
    }
    return labels;
}

function generateRandomData(count, min, max) {
    const data = [];
    for (let i = 0; i < count; i++) {
        data.push(Math.floor(Math.random() * (max - min + 1)) + min);
    }
    return data;
}

// Modal Functions
function openDeviceModal(device) {
    const modal = document.getElementById('deviceModal');
    modal.classList.add('show');

    // Update modal content with device data
    document.getElementById('modalDeviceName').textContent = `${device.name} - Details`;
    document.getElementById('modalDeviceId').textContent = device.id;
    document.getElementById('modalDeviceType').textContent = device.type;
    document.getElementById('modalDeviceModel').textContent = device.model;
    document.getElementById('modalDeviceSerial').textContent = device.serial;
    document.getElementById('modalDeviceFirmware').textContent = device.firmware;
    document.getElementById('modalDeviceIP').textContent = device.ip;
    document.getElementById('modalDeviceMAC').textContent = device.mac;
    document.getElementById('modalDeviceSubnet').textContent = device.subnet;
    document.getElementById('modalDeviceGateway').textContent = device.gateway;
    document.getElementById('modalDeviceVLAN').textContent = device.vlan;

    // Initialize performance chart in modal
    initializeDevicePerformanceChart();

    // Reset tabs to show performance by default
    switchTab('performance');
}

function closeDeviceModal() {
    const modal = document.getElementById('deviceModal');
    modal.classList.remove('show');
}

function switchTab(tabName) {
    // Update tab buttons
    const tabButtons = document.querySelectorAll('.tab-button');
    tabButtons.forEach(btn => {
        btn.classList.remove('active');
        if (btn.textContent.toLowerCase().includes(tabName.toLowerCase()) ||
            (tabName === 'performance' && btn.textContent.includes('Performance')) ||
            (tabName === 'ports' && btn.textContent.includes('Ports')) ||
            (tabName === 'logs' && btn.textContent.includes('Logs')) ||
            (tabName === 'config' && btn.textContent.includes('Configuration'))) {
            btn.classList.add('active');
        }
    });

    // Update tab panes
    const tabPanes = document.querySelectorAll('.tab-pane');
    tabPanes.forEach(pane => {
        pane.classList.remove('active');
    });

    const activePane = document.getElementById(`${tabName}-tab`);
    if (activePane) {
        activePane.classList.add('active');
    }
}

function initializeDevicePerformanceChart() {
    const ctx = document.getElementById('devicePerformanceChart');
    if (!ctx) return;

    // Clear any existing chart
    if (window.devicePerfChart) {
        window.devicePerfChart.destroy();
    }

    // Generate static data once
    const cpuData = [45, 48, 52, 47, 49, 46, 51, 48, 45, 47, 50, 48];
    const memoryData = [62, 64, 61, 63, 65, 62, 60, 63, 64, 62, 61, 62];

    window.devicePerfChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: ['11:00', '10:00', '9:00', '8:00', '7:00', '6:00', '5:00', '4:00', '3:00', '2:00', '1:00', '0:00'],
            datasets: [{
                label: 'CPU Usage (%)',
                data: cpuData,
                borderColor: '#3b82f6',
                backgroundColor: 'rgba(59, 130, 246, 0.1)',
                tension: 0.3
            }, {
                label: 'Memory Usage (%)',
                data: memoryData,
                borderColor: '#10b981',
                backgroundColor: 'rgba(16, 185, 129, 0.1)',
                tension: 0.3
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            animation: false,
            plugins: {
                legend: {
                    position: 'top',
                }
            },
            scales: {
                y: {
                    beginAtZero: true,
                    max: 100
                }
            }
        }
    });
}

// Device action functions
function configureDevice(deviceId) {
    showNotification(`Opening configuration for device ${deviceId}`);
}

function monitorDevice(deviceId) {
    showNotification(`Opening monitoring for device ${deviceId}`);
}

function rebootDevice(deviceId) {
    if (confirm(`Are you sure you want to reboot device ${deviceId}?`)) {
        showNotification(`Rebooting device ${deviceId}...`);
    }
}

function openDeviceConsole() {
    showNotification('Opening device console...');
}

function restartDevice() {
    if (confirm('Are you sure you want to restart this device?')) {
        showNotification('Device restart initiated');
    }
}

function shutdownDevice() {
    if (confirm('Are you sure you want to shutdown this device? This action cannot be undone.')) {
        showNotification('Device shutdown initiated');
    }
}

function pingDevice() {
    showNotification('Pinging device... Response time: 12ms');
}

function backupConfig() {
    showNotification('Configuration backup started...');
}

function restoreConfig() {
    showNotification('Opening configuration restore dialog...');
}

function editConfig() {
    showNotification('Opening configuration editor...');
}

// Close modal when clicking outside
window.onclick = function(event) {
    const modal = document.getElementById('deviceModal');
    if (event.target === modal) {
        closeDeviceModal();
    }
}

// Add ESC key to close modal
document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        closeDeviceModal();
    }
});

const style = document.createElement('style');
style.textContent = `
    @keyframes slideIn {
        from { transform: translateX(100%); opacity: 0; }
        to { transform: translateX(0); opacity: 1; }
    }

    @keyframes slideOut {
        from { transform: translateX(0); opacity: 1; }
        to { transform: translateX(100%); opacity: 0; }
    }

    .fa-spin {
        animation: spin 1s linear infinite;
    }

    @keyframes spin {
        from { transform: rotate(0deg); }
        to { transform: rotate(360deg); }
    }
`;
document.head.appendChild(style);