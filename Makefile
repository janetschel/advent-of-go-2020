day = 3

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir calendar/day-0$(day); \
  		cp template-day.go calendar/day-0$(day)/day0$(day).go; \
  		cp template-day.go calendar/day-0$(day)/day0$(day)_pt02.go; \
  	else \
  		mkdir calendar/day-$(day); \
		cp template-day.go calendar/day-$(day)/day$(day).go; \
		cp template-day.go calendar/day-$(day)/day$(day)_pt02.go; \
    fi
	@echo "Files successfully created.. happy hacking :)"
	@echo "INFO: puzzle input still needs to be fetched"
	@git add calendar/


