import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CalendarComponent } from './calendar.component';

describe('CalendarComponent', () => {
  let component: CalendarComponent;
  let fixture: ComponentFixture<CalendarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CalendarComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CalendarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    //expect(component).toBeTruthy();
  });

  it('should initialize the dataSource property with data', () => {
    expect(component.dataSource.length).greaterThan(0);
  });

  it('should initialize the currentTimeIndicator property', () => {
    expect(component.currentTimeIndicator).equal(true);
  });

  it('should initialize the shadeUntilCurrentTime property', () => {
    expect(component.shadeUntilCurrentTime).equal(true);
  });

  it('should initialize the view property', () => {
    expect(component.view).equal('day');
  });

  it('should initialize the views property', () => {
    expect(component.views).equal(['day', 'week', 'month', 'timelineDay', 'timelineWeek', 'timelineMonth']);
  });

  it('should initialize the firstDayOfWeek property', () => {
    expect(component.firstDayOfWeek).equal(1);
  });

  it('should initialize the dataSource property with specific data', () => {
    const data = component.getData();
    expect(data.length).equal(2);
    expect(data[0].label).equal('Example Event');
    expect(data[1].label).equal('Jenna and Kenneth 21st Birthday');
  });
});
