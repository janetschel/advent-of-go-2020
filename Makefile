day = $(shell date +'%-d')

new:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir calendar/day-0$(day); \
  		cp template calendar/day-0$(day)/day0$(day).go; \
  		cp template calendar/day-0$(day)/day0$(day)_pt02.go; \
  		touch calendar/day-$0(day)/md_day0$(day).md; \
  	else \
  		mkdir calendar/day-$(day); \
		cp template calendar/day-$(day)/day$(day).go; \
		cp template calendar/day-$(day)/day$(day)_pt02.go; \
		touch calendar/day-$(day)/md_day$(day).md; \
    fi
	@echo "Files successfully created.. happy hacking :)"
	@echo "INFO: puzzle input still needs to be fetched"
	@git add calendar/


