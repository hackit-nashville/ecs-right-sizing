import pandas as pd
import sys
import json

def calculate(data, interval):
    with open(data) as json_file:
        temp = json.load(json_file)

    new_dict = temp['MetricDataResults']
    df = pd.DataFrame.from_dict(new_dict)
    timestamps = df.Timestamps[0]
    task_count = df.Values[0]
    memory_reserved = df.Values[1]
    memory_utilized = df.Values[2]

    data = pd.DataFrame(
        {'Timestamps': timestamps,
         'RunningTaskCount': task_count,
         'MemoryReserved': memory_reserved,
         'MemoryUtilized': memory_utilized
         }
    )

    data['Timestamps'] = pd.to_datetime(data['Timestamps'])

    new_df = data.drop(columns=['RunningTaskCount', 'MemoryReserved'])

    if interval.upper() == 'HOUR':

        new_df['Hour'] = new_df['Timestamps'].apply(lambda x: x.hour)
        final_df = new_df.groupby(['Hour'])['MemoryUtilized'].mean().reset_index()
        final_df.to_csv(path_or_buf='hour.csv', sep=',', index=False)
        with pd.option_context('display.max_rows', None, 'display.max_columns', None):
            print(final_df)
        return

    elif interval.upper() == 'DOW':

        new_df['DOW'] = new_df['Timestamps'].apply(lambda x: x.dayofweek)
        final_df = new_df.groupby(['DOW'])['MemoryUtilized'].mean().reset_index()
        final_df.to_csv(path_or_buf='dow.csv', sep=',', index=False)
        with pd.option_context('display.max_rows', None, 'display.max_columns', None):
            print(final_df)
        return

    elif interval.upper() == 'DOW_HOURLY':

        new_df['Hour'] = new_df['Timestamps'].apply(lambda x: x.hour)
        new_df['DOW'] = new_df['Timestamps'].apply(lambda x: x.dayofweek)
        final_df = new_df.groupby(['DOW', ' Hour'])['MemoryUtilized'].mean().reset_index()
        final_df.to_csv(path_or_buf='dow_hourly.csv', sep=',', index=False)
        with pd.option_context('display.max_rows', None, 'display.max_columns', None):
            print(final_df)
        return

    else:
        print("Not a valid interval. Please use: HOUR, DOW, or DOW_HOURLY")
        return


if __name__ == '__main__':
    calculate(*sys.argv[1:])
