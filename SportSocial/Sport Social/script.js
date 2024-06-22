document.addEventListener('DOMContentLoaded', () => {
    const saveButton = document.querySelector('.save-button');
    const locationDropdownButton = document.querySelector('.custom-button');
    const locationDropdownContent = document.querySelector('.dropdown-content');
    const activityDropdownButton = document.querySelector('.custom-button2');
    const activityDropdownContent = document.querySelector('.dropdown-content2');

    locationDropdownButton.addEventListener('click', () => {
        locationDropdownContent.classList.toggle('show');
    });

    activityDropdownButton.addEventListener('click', () => {
        activityDropdownContent.classList.toggle('show');
    });

    saveButton.addEventListener('click', () => {
        const locationCheckboxes = document.querySelectorAll('.dropdown-content input[type="checkbox"]');
        const activityCheckboxes = document.querySelectorAll('.dropdown-content2 input[type="checkbox"]');
        const selectedLocations = [];
        const selectedActivities = [];

        locationCheckboxes.forEach(checkbox => {
            if (checkbox.checked) {
                selectedLocations.push(checkbox.name);
            }
        });

        activityCheckboxes.forEach(checkbox => {
            if (checkbox.checked) {
                selectedActivities.push(checkbox.name);
            }
        });

        alert('Selected locations: ' + selectedLocations.join(', ') + '\nSelected activities: ' + selectedActivities.join(', '));
    });
    
});
